package main

import (
	"fmt"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	e.POST("/login", func(c echo.Context) error {
		c.String(500, "db error")
		return nil
	})

	e.Use(
		middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
			LogError:     true,
			LogStatus:    true,
			LogRequestID: true,
			LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
				if v.Status < 400 {
					return nil
				}

				in, _ := c.FormParams()
				out := v.FormValues

				c.Logger().Printj(log.JSON{
					"s":   v.Status,
					"err": v.Error,
					"id":  v.RequestID,
					"in":  in,
					"out": out,
				})
				return nil
			},
		}),
	)

	err := e.Start(":80")
	if err != nil {
		fmt.Printf("err server %v", err)
	}
}

// curl -X POST 'http://localhost/login' -d "login=a.sulaev&password=123"
