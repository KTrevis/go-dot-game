package main

import (
	"fmt"
	"net/http"
	"server/Client"
	"server/database"
	"server/views"
	"server/views/api"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

func createWebsocket(context *gin.Context, manager *client.WebSocketManager) {
	socket, err := upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	manager.AddClient(socket)
	go manager.Clients[socket].Loop()
}

func setupViews(router *gin.Engine, manager *client.WebSocketManager) {

	router.GET("/", views.Index)

	router.GET("/websocket", func(context *gin.Context) {
		createWebsocket(context, manager)
	})

	router.POST("/api/register", func(c *gin.Context) {
		api.Register(c, manager.DB)
	})

	router.Static("/game", "./game")
	router.LoadHTMLGlob("./templates/*")
}

func main() {
	manager := client.NewWebSocketManager()
	manager.DB = database.SetupDB()
	router := gin.Default()
	setupViews(router, manager)
	router.Run(":80")
}
