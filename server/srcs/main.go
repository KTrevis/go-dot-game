package main

import (
	"server/srcs/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var manager = NewWebSocketManager()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createWebsocket(context *gin.Context) {
	socket, _ := upgrader.Upgrade(context.Writer, context.Request, nil)
	manager.AddClient(socket)
	go manager.Clients[socket].Loop()
}

func main() {
	manager.DB = database.Setup()
	router := gin.Default()
	router.GET("/websocket", createWebsocket)
	router.Run(":8080")
}
