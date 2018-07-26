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
	Server ServerConfig `json:"server"`
	Som    SomConfig    `json:"db"`
}
type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type SomConfig struct {
	Size int `json:"size"`
}

func main() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)

	err := som.InitMapByEuclidean(config.Som.Size)
	if err != nil {
		panic(err)
	}

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア（ログとか）はここ
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MainPage())

	// サーバー起動
	e.Start(":" + config.Server.Port) //ポート番号指定してね
}
