package handler

import (
	"net/http"

	"../som"
	"github.com/labstack/echo"
)

// Request格納用
type traitAPIRequest struct {
	unit som.Unit
}

// Response格納用
type traitAPIResponse struct {
	distance float64
}
type mapAPIResponse struct {
	sMap [][]som.Unit
}
type errorAPIResponse struct {
	message string
}

// MainPage サンプルページ
func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Hello World")
	}
}

// TraitAPI  学習API
func TraitAPI(tc chan som.TraitChan) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(traitAPIRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		var data som.TraitChan
		data.Unit = req.unit
		data.ResDistance = make(chan float64)
		tc <- data
		value, ok := <-data.ResDistance
		if !ok {
			return c.JSON(http.StatusOK, errorAPIResponse{"trait failed"})
		}
		return c.JSON(http.StatusOK, traitAPIResponse{value})
	}
}

// MapAPI  学習MAP取得API
func MapAPI(mc chan som.MapChan) echo.HandlerFunc {
	return func(c echo.Context) error {
		var data som.MapChan
		data.ResMap = make(chan [][]som.Unit)
		mc <- data

		value, ok := <-data.ResMap
		if !ok {
			return c.JSON(http.StatusOK, errorAPIResponse{"get map Faild"})
		}
		return c.JSON(http.StatusOK, mapAPIResponse{value})
	}
}
