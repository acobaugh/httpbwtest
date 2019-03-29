package main

import (
	"fmt"
	"net/http"

	"github.com/dustin/go-humanize"
	"github.com/labstack/echo"
	elm "github.com/pennstate/echo-logrusmiddleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"mcquay.me/trash"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	e := echo.New()
	e.Logger = elm.Logger{logrus.StandardLogger()}
	e.Use(elm.Hook())

	e.GET("/:pattern", getData)
	e.GET("/", getData)
	e.POST("/", postData)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func getData(c echo.Context) error {
	var size uint64
	var err error

	querySize := c.QueryParam("size")
	if querySize != "" {
		size, err = humanize.ParseBytes(querySize)
		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Bad value for size: %s", err))
		}
	} else {
		size = 1
	}

	var r io.Reader
	switch c.Param("pattern") {
	case "fs":
		r = trash.Fs
	case "hilo":
		r = trash.HiLo
	case "lohi":
		r = trash.LoHi
	case "zeros":
		r = trash.Zeros
	case "random":
		r = trash.Random
	default:
		r = trash.Reader
	}

	lr := &io.LimitedReader{
		R: r,
		N: int64(size),
	}
	return c.Stream(http.StatusOK, "application/octet-stream", lr)
}
func postData(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
