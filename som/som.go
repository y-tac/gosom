package som

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// MaxValue som unitの最大値
const MaxValue = 256

// Unit 要素型
type Unit struct {
	Red   int
	Blue  int
	Green int
}

// Factor 要素係数型
type Factor struct {
	Red   float64
	Blue  float64
	Green float64
}

// Som Unitで構成された球面SOM
type Som struct {
	// マップ
	sMap [][]Unit
	// 中点
	midpointX int
	midpointY int
	// 近傍半径
	radius int
	// 不偏分散係数
	uVariance Factor
}

// DataMap データ実体
var DataMap Som

// distanceFunc
var distanceFunc func(Unit, Unit) int

// euclideandistance デフォルト距離計測関数。a,b間のユークリッド距離の二乗を計算する
func euclideandSqDistance(a Unit, b Unit) int {
	return (((a.Red - b.Red) * (a.Red - b.Red)) + ((a.Blue - b.Blue) * (a.Blue - b.Blue)) + ((a.Green - b.Green) * (a.Green - b.Green)))
}

// InitMapByEuclidean SOM初期化関数(ユークリッド距離の二乗で計算)
func initMapByEuclidean(r int) error {
	return initMap(r, euclideandSqDistance)
}

// InitMap SOM初期化関数
func initMap(r int, fn func(Unit, Unit) int) error {
	distanceFunc = fn
	rand.Seed(time.Now().UnixNano())
	// 中点の初期化
	DataMap.midpointX = r / 2
	DataMap.midpointY = r / 2

	// Mapの初期化
	DataMap.sMap = make([][]Unit, r)
	for x := 0; x < r; x++ {
		DataMap.sMap[x] = make([]Unit, r)
		for y := 0; y < r; y++ {
			DataMap.sMap[x][y].Red = rand.Intn(MaxValue)
			DataMap.sMap[x][y].Blue = rand.Intn(MaxValue)
			DataMap.sMap[x][y].Green = rand.Intn(MaxValue)
		}
	}

	// 近傍半径の初期化
	DataMap.uVariance = calcUVariance(DataMap.sMap)
	DataMap.radius = calcRadius(DataMap.uVariance, r)

	return nil
}

// som index変換
func getRadiusIndex(i int) int {
	return (i + len(DataMap.sMap)) % (len(DataMap.sMap))
}

// 不偏分散計算関数
func calcUVariance(values [][]Unit) (resUVariance Factor) {
	//平均値の計算
	var RedAve, BlueAve, GreenAve float64
	num := 0
	for x := 0; x < len(values); x++ {
		for y := 0; y < len(values[x]); y++ {
			RedAve += float64(values[x][y].Red)
			BlueAve += float64(values[x][y].Blue)
			GreenAve += float64(values[x][y].Green)
			resUVariance.Red += float64(values[x][y].Red * values[x][y].Red)
			resUVariance.Blue += float64(values[x][y].Blue * values[x][y].Blue)
			resUVariance.Green += float64(values[x][y].Green * values[x][y].Green)
			num++
		}
	}
	fnum := float64(num)
	RedAve = RedAve / fnum
	BlueAve = BlueAve / fnum
	GreenAve = GreenAve / fnum

	resUVariance.Red = (resUVariance.Red - fnum*RedAve) / (fnum - 1)
	resUVariance.Blue = (resUVariance.Blue - fnum*BlueAve) / (fnum - 1)
	resUVariance.Green = (resUVariance.Green - fnum*GreenAve) / (fnum - 1)
	return
}

// 近傍半径更新関数
func calcRadius(uv Factor, r int) (result int) {
	vectorX, vectorY, vectorZ := math.Sqrt(uv.Red), math.Sqrt(uv.Blue), math.Sqrt(uv.Green)
	return int(float64(r) * math.Sqrt(vectorX*vectorX+vectorY*vectorY+vectorZ*vectorZ) / math.Sqrt(3*MaxValue*MaxValue))
}

// 係数計算関数
func calcRFactor(vectorX int, vectorY int, uVariance Factor, r int) (resFactor Factor) {
	sqScala := float64(vectorX*vectorX + vectorY*vectorY)
	scala := math.Sqrt(sqScala)
	resFactor.Red = math.Exp(-sqScala/(2*uVariance.Red)) * scala / float64(r)
	resFactor.Blue = math.Exp(-sqScala/(2*uVariance.Blue)) * scala / float64(r)
	resFactor.Green = math.Exp(-sqScala/(2*uVariance.Green)) * scala / float64(r)
	return
}

// Unit更新関数
func updateUnit(before Unit, t Unit, rFactor Factor) (r Unit) {
	r = before
	r.Red += int(rFactor.Red * float64(t.Red-before.Red))
	r.Green += int(rFactor.Blue * float64(t.Green-before.Green))
	r.Blue += int(rFactor.Green * float64(t.Blue-before.Blue))
	return r
}

// Trait SOM更新関数。Unitのパラメータを受け取って中点からの距離を返却する
func trait(t Unit) float64 {
	// 初期値として適当な場所の距離を入れておく。
	BMUindexX, BMUindexY := 0, 0
	var dMin = distanceFunc(t, DataMap.sMap[0][0])

	// BMU探索処理
	for x := 0; x < len(DataMap.sMap); x++ {
		for y := 0; y < len(DataMap.sMap[x]); y++ {
			d := distanceFunc(t, DataMap.sMap[x][y])
			if d < dMin {
				dMin = d
				BMUindexX, BMUindexY = x, y
			}
		}
	}

	// 更新処理
	for x := BMUindexX - DataMap.radius; x < (BMUindexX + DataMap.radius); x++ {
		for y := BMUindexY - DataMap.radius; y < (BMUindexY + DataMap.radius); y++ {
			indexX := getRadiusIndex(x)
			indexY := getRadiusIndex(y)
			rFactor := calcRFactor(DataMap.midpointX-x, DataMap.midpointY-y, DataMap.uVariance, len(DataMap.sMap))
			DataMap.sMap[indexX][indexY] = updateUnit(DataMap.sMap[indexX][indexY], t, rFactor)
		}
	}
	// 標準偏差を更新する
	DataMap.uVariance = calcUVariance(DataMap.sMap)
	// 近傍半径を更新する
	DataMap.radius = calcRadius(DataMap.uVariance, len(DataMap.sMap))
	fmt.Println("BMU:", BMUindexX, BMUindexY, "MID:", DataMap.midpointX, DataMap.midpointY, DataMap.radius, DataMap.uVariance)

	// 中点からの距離を計算する
	resDist := math.Sqrt(math.Pow(float64(BMUindexX-DataMap.midpointX), 2) + math.Pow(float64(BMUindexY-DataMap.midpointY), 2))
	// 中点を更新する
	if DataMap.midpointX < BMUindexX {
		DataMap.midpointX++
	} else {
		DataMap.midpointX--
	}
	if DataMap.midpointY < BMUindexY {
		DataMap.midpointY++
	} else {
		DataMap.midpointY--
	}
	return resDist
}
