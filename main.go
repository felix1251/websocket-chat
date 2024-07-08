package main

import (
	"fmt"

	"chat/router"
	"chat/ws"
)

func main() {
	fmt.Println("start websocket server..")

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)

	fmt.Println("hub created...")

	router.Init(wsHandler)
	router.Start("0.0.0.0:3005")
}
