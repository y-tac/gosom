package main

import (
	"encoding/json"
	"io/ioutil"

	"./handler"
	"./som"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Config struct {
	Server ServerConfig  `json:"server"`
	Som    som.SomConfig `json:"db"`
}
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
	chset := som.SomRoutine(config.Som)

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())
	e.GET("/trait", handler.TraitAPI(chset.TraitCh))
	e.GET("/map", handler.MapAPI(chset.MapCh))

	// サーバー起動
	e.Start(":" + config.Server.Port)
}
