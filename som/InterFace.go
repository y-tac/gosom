package som

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

// Config SOMが要求するコンフィグ
type Config struct {
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

// Routine SOMスレッド処理関数
func Routine(conf Config, gosomDistance prometheus.Gauge) (chset ChanSet) {
	chset.TraitCh = make(chan TraitChan)
	chset.MapCh = make(chan MapChan)

	go func(chset ChanSet, conf Config, gosomDistance prometheus.Gauge) {
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
				fmt.Println(traitCh.Unit)
				distance := trait(traitCh.Unit)
				gosomDistance.Set(distance)
				traitCh.ResDistance <- distance
			case mapCh, ok := <-chset.MapCh:
				if !ok {
					return
				}
				mapCh.ResMap <- mapgenerate()
			}
		}
	}(chset, conf, gosomDistance)
	return
}
