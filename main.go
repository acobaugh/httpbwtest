package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	elm "github.com/pennstate/echo-logrusmiddleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("port", 8080)
	port := viper.GetInt("port")

	e := echo.New()
	e.Logger = elm.Logger{logrus.StandardLogger()}
	e.Use(elm.Hook())

	e.GET("/", recvHandler)
	e.POST("/", sendHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func recvHandler(c echo.Context) error {
	size := c.QueryParam("size")

	return c.String(http.StatusOK, fmt.Sprintf("this is a test. size = %v", size))
}
func sendHandler(c echo.Context) error {
	return nil
}
