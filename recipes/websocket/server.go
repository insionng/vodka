package main

import (
	"fmt"

	"github.com/insionng/vodka"
	"github.com/insionng/vodka/libraries/net/websocket"
	mw "github.com/insionng/vodka/middleware"
)

func main() {
	e := vodka.New()

	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Static("/", "public")
	e.WebSocket("/ws", func(c *vodka.Context) (err error) {
		ws := c.Socket()
		msg := ""

		for {
			if err = websocket.Message.Send(ws, "Hello, Client!"); err != nil {
				return
			}
			if err = websocket.Message.Receive(ws, &msg); err != nil {
				return
			}
			fmt.Println(msg)
		}
		return
	})

	e.Run(":9000")
}
