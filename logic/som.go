package som

import (
	"math"
	"math/rand"
	"time"
)

const maxValue = 256

// Unit 要素型
type Unit struct {
	red   int
	blue  int
	green int
}

// Factor 要素係数型
type Factor struct {
	red   float64
	blue  float64
	green float64
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
	return (((a.red - b.red) * (a.red - b.red)) + ((a.blue - b.blue) * (a.blue - b.blue)) + ((a.green - b.green) * (a.green - b.green)))
}

// InitMapByEuclidean SOM初期化関数(ユークリッド距離の二乗で計算)
func InitMapByEuclidean(r int) bool {
	return InitMap(r, euclideandSqDistance)
}

// InitMap SOM初期化関数
func InitMap(r int, fn func(Unit, Unit) int) bool {
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
			DataMap.sMap[x][y].red = rand.Intn(maxValue)
			DataMap.sMap[x][y].blue = rand.Intn(maxValue)
			DataMap.sMap[x][y].green = rand.Intn(maxValue)
		}
	}

	// 近傍半径の初期化
	DataMap.uVariance = calcUVariance(DataMap.sMap)
	DataMap.radius = calcRadius(DataMap.uVariance, r)

	return true
}

// som index変換
func getRadiusIndex(i int) int {
	return i % (len(DataMap.sMap))
}

// 不偏分散計算関数
func calcUVariance(values [][]Unit) (resUVariance Factor) {
	//平均値の計算
	var redAve, blueAve, greenAve float64
	num := 0
	for x := 0; x < len(values); x++ {
		for y := 0; y < len(values[x]); y++ {
			redAve += float64(values[x][y].red)
			blueAve += float64(values[x][y].blue)
			greenAve += float64(values[x][y].green)
			resUVariance.red += float64(values[x][y].red * values[x][y].red)
			resUVariance.blue += float64(values[x][y].blue * values[x][y].blue)
			resUVariance.green += float64(values[x][y].green * values[x][y].green)
			num++
		}
	}
	fnum := float64(num)
	redAve = redAve / fnum
	blueAve = blueAve / fnum
	greenAve = greenAve / fnum

	resUVariance.red = (resUVariance.red - fnum*redAve) / (fnum - 1)
	resUVariance.blue = (resUVariance.blue - fnum*blueAve) / (fnum - 1)
	resUVariance.green = (resUVariance.green - fnum*greenAve) / (fnum - 1)
	return
}

// 近傍半径更新関数
func calcRadius(uv Factor, r int) (result int) {
	vectorX, vectorY, vectorZ := math.Sqrt(uv.red), math.Sqrt(uv.blue), math.Sqrt(uv.green)
	return int(float64(r) * math.Sqrt(vectorX*vectorX+vectorY*vectorY+vectorZ*vectorZ) / math.Sqrt(3*maxValue*maxValue))
}

// 係数計算関数
func calcRFactor(vectorX int, vectorY int, uVariance Factor) (resFactor Factor) {
	scala := float64(vectorX*vectorX + vectorY*vectorY)
	resFactor.red = math.Exp(-scala / (2 * uVariance.red))
	resFactor.blue = math.Exp(-scala / (2 * uVariance.blue))
	resFactor.green = math.Exp(-scala / (2 * uVariance.green))
	return
}

// Unit更新関数
func updateUnit(before Unit, t Unit, rFactor Factor) (r Unit) {
	r = before
	r.red += int(rFactor.red * float64(t.red-before.red))
	r.green += int(rFactor.blue * float64(t.green-before.green))
	r.blue += int(rFactor.green * float64(t.blue-before.blue))
	return r
}

// Trait SOM更新関数。Unitのパラメータを受け取って中点からの距離を返却する
func Trait(t Unit) float64 {
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
			rFactor := calcRFactor(DataMap.midpointX-x, DataMap.midpointY-y, DataMap.uVariance)
			DataMap.sMap[indexX][indexY] = updateUnit(DataMap.sMap[indexX][indexY], t, rFactor)
		}
	}
	// 標準偏差を更新する
	DataMap.uVariance = calcUVariance(DataMap.sMap)
	// 近傍半径を更新する
	DataMap.radius = calcRadius(DataMap.uVariance, len(DataMap.sMap))
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
	// 中点からの距離を計算する
	resDist := math.Sqrt(math.Pow(float64(BMUindexX-DataMap.midpointX), 2) + math.Pow(float64(BMUindexY-DataMap.midpointY), 2))
	return resDist
}

// Map SOM取得関数
func Map() [][]Unit {
	return DataMap.sMap
}
