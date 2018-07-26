package som

// SomConfig SOMが要求するコンフィグ
type SomConfig struct {
	Size int `json:"size"`
}

// SomChan SOM Routineで利用するチャネル
type SomChan struct {
	Unit Unit
	DRes chan float64
}

// SomRoutine SOMスレッド処理関数
func SomRoutine(conf SomConfig) (sc chan SomChan) {
	sc = make(chan SomChan)
	go func(sc chan SomChan) {
		err := initMapByEuclidean(conf.Size)
		if err != nil {
			panic(err)
		}
		for {
			ch, ok := <-sc
			if !ok {
				return
			}
			ch.DRes <- trait(ch.Unit)
		}
	}(sc)
	return
}

// Map SOM取得関数
func Map() [][]Unit {
	return DataMap.sMap
}
