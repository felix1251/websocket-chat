package router

import (
	"chat/ws"

	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func Init(wsHandler *ws.Handler) {
	r = gin.Default()

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
}

func Start(addr string) error {
	return r.Run(addr)
}
