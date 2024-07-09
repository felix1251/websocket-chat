package main

import (
	"chat/router"
	"chat/ws"
)

func main() {
	// Initialize WebSocket hub and handler
	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	// Initialize router and start the server
	router.Init(wsHandler)
	router.Start("0.0.0.0:8080")
}
