package som

import (
	"math/rand"
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport

	"github.com/prometheus/client_golang/prometheus"
)

func TestInitToTrait(t *testing.T) {
	conf := SomConfig{Size: 100}
	traitNum := int(50)
	err := initMapByEuclidean(conf.Size)
	if err != nil {
		t.Fatalf("failed initMapByEuclidean %#v", err)
	}
	for i := 0; i < traitNum; i++ {
		unit := Unit{rand.Intn(MaxValue), rand.Intn(MaxValue), rand.Intn(MaxValue)}
		dist := trait(unit)
		if dist < 0 {
			t.Fatalf("failed calc distance %#v", dist)
		}
	}

	maps := mapgenerate()
	if len(maps) != conf.Size {
		t.Fatalf("failed map data %#v", maps)
	}
}

func TestSomroutine(t *testing.T) {
	conf := SomConfig{Size: 100}
	gosomDistance := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gosom",
		Name:      "distance",
		Help:      "Distance of midpoint to traitpoint",
	})
	set := SomRoutine(conf, gosomDistance)
	traitNum := int(50)
	for i := 0; i < traitNum; i++ {
		var data TraitChan
		data.Unit = Unit{rand.Intn(MaxValue), rand.Intn(MaxValue), rand.Intn(MaxValue)}
		data.ResDistance = make(chan float64)

		set.TraitCh <- data
		value, ok := <-data.ResDistance
		if !ok {
			t.Fatalf("failed channel %#v", data)
		}
		if value < 0 {
			t.Fatalf("failed calc distance %#v", value)
		}

	}

	var mapc MapChan
	mapc.ResMap = make(chan [][]Unit)
	set.MapCh <- mapc

	value, ok := <-mapc.ResMap
	if !ok {
		t.Fatalf("failed test %#v", mapc)
	}
	if len(value) != conf.Size {
		t.Fatalf("failed map data %#v", value)
	}

}
