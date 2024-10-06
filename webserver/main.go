package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {

			// Read
			msg := ""
			err := websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Fatal(err)
			}
			fmt.Printf("%s\n", msg)
			if msg != "" {
				err = websocket.Message.Send(ws, msg)
				if err != nil {
					c.Logger().Error(err)
				}
				msg = ""
			}

		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/", "../public")
	e.GET("/ws2", hello)
	e.Logger.Fatal(e.Start(":1323"))
}
