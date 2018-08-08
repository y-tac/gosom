package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/y-tac/gosom/dataporter"
	"github.com/y-tac/gosom/handler"
	"github.com/y-tac/gosom/som"
)

// Config 構造体。サブパッケージのconfigも設定する
type Config struct {
	Server     ServerConfig      `json:"server"`
	Som        som.Config        `json:"som"`
	DataPorter dataporter.Config `json:"dataporter"`
}

// ServerConfig サーバconfigを設定
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	fmt.Println(config)

	gosomDistance := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "gosom",
		Name:      "distance",
		Help:      "Distance of midpoint to traitpoint",
	})
	prometheus.MustRegister(gosomDistance)

	dataporterQuit := make(chan bool)
	go dataporter.Dataporter(config.DataPorter, dataporterQuit)
	defer closeGoroutine(dataporterQuit)
	gosomQuit := make(chan bool)
	chset := som.MakeChannelRoutine()
	go som.Routine(chset, config.Som, gosomDistance, gosomQuit)
	defer closeGoroutine(gosomQuit)
	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/front", "front/dist")

	// ルーティング
	e.GET("/hello", handler.MainPage())
	e.POST("/trait", handler.TraitAPI(chset.TraitCh))
	e.GET("/map", handler.MapAPI(chset.MapCh))
	e.GET("/metrics", echo.WrapHandler(prometheus.Handler()))
	// サーバー起動
	e.Start(":" + config.Server.Port)
}

// goroutine終了関数
func closeGoroutine(quit chan bool) {
	quit <- true
	return
}
