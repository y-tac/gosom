package som

// SomConfig SOMが要求するコンフィグ
type SomConfig struct {
	Size int `json:"size"`
}

// ChanSet SOM通信用チャンネル
type ChanSet struct {
	TraitCh chan TraitChan
	MapCh   chan MapChan
}

// TraitChan 学習情報やり取り用
type TraitChan struct {
	Unit        Unit
	ResDistance chan float64
}

// MapChan SOM Routineで利用するチャネル
type MapChan struct {
	ResMap chan [][]Unit
}

// SomRoutine SOMスレッド処理関数
func SomRoutine(conf SomConfig) (chset ChanSet) {
	chset.TraitCh = make(chan TraitChan)
	chset.MapCh = make(chan MapChan)

	go func(chset ChanSet, conf SomConfig) {
		err := initMapByEuclidean(conf.Size)
		if err != nil {
			panic(err)
		}
		for {
			select {
			case traitCh, ok := <-chset.TraitCh:
				if !ok {
					return
				}
				traitCh.ResDistance <- trait(traitCh.Unit)
			case mapCh, ok := <-chset.MapCh:
				if !ok {
					return
				}
				mapCh.ResMap <- DataMap.sMap
			}
		}
	}(chset, conf)
	return
}
