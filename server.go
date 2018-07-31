package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/y-tac/gosom/dataporter"
	"github.com/y-tac/gosom/handler"
	"github.com/y-tac/gosom/som"
)

// Config 構造体。サブパッケージのconfigも設定する
type Config struct {
	Server     ServerConfig                `json:"server"`
	Som        som.SomConfig               `json:"som"`
	DataPorter dataporter.DataPorterConfig `json:"dataporter"`
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
	go dataporter.Dataporter(config.DataPorter)
	chset := som.SomRoutine(config.Som)

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

	// サーバー起動
	e.Start(":" + config.Server.Port)
}
