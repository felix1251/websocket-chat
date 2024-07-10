package router

import (
	"chat/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init(wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/ws/create_room", wsHandler.CreateRoom)
	r.GET("/ws/join_room/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/get_rooms", wsHandler.GetRooms)
	r.GET("/ws/get_clients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
