package handler

import (
	"net/http"

	"../som"
	"github.com/labstack/echo"
)

// MainPage サンプルページ
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Hello World")
	}
}

// TraitAPI  学習API
func TraitAPI(sc chan som.SomChan) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data som.SomChan
		data.Unit.Red = 0
		data.Unit.Blue = 0
		data.Unit.Green = 0
		sc <- data
		data, ok := <-sc
		if !ok {

		}
		return c.String(http.StatusOK, "Hello World")
	}
}
