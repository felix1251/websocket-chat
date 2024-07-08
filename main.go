package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func websocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading connection:", err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("Recieve message: %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)

		if err != nil {
			log.Println(err)
			break
		}
	}
}

func main() {
	fmt.Println("start websocket server..")

	http.HandleFunc("/ws", websocketHandler)

	log.Fatal(http.ListenAndServe(":3005", nil))
}
