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

// Som Unitで構成された球面SOM
type Som struct {
	// マップ
	sMap [][]Unit
	// 中点
	midpointX int
	midpointY int
	// 近傍半径
	radius int
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
	// 近傍半径の初期化
	DataMap.radius = r / 2
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

	return true
}

// som index変換
func getRadiusIndex(i int) int {
	return i % (len(DataMap.sMap))
}

// 近傍半径更新関数
func updateRadius() int {
	return len(DataMap.sMap) / 2
}

// 係数計算
func calcRFactor() float64 {
	return 1.0
}

// Unit更新関数
func updateUnit(before Unit, t Unit, rFactor float64) (r Unit) {
	r = before
	r.red += int(rFactor * float64(t.red-before.red))
	r.green += int(rFactor * float64(t.green-before.green))
	r.blue += int(rFactor * float64(t.blue-before.blue))
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
			rFactor := calcRFactor()
			DataMap.sMap[indexX][indexY] = updateUnit(DataMap.sMap[indexX][indexY], t, rFactor)
		}
	}
	// 近傍半径を更新する
	DataMap.radius = updateRadius()
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
