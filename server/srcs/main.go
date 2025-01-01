package main

import (
	"fmt"
	"log"
	"net/http"
	"server/CLI"
	"server/Client"
	"server/database"
	gamemaps "server/maps"
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

func websocketCLI(context *gin.Context, manager *client.WebSocketManager) {
	socket, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Printf("createdWebsocket failed: %s", err.Error())
		return
	}

	cli := &cli.CLI {
		Socket: socket,
		Manager: manager,
	}
	go cli.Loop()
}

func websocketClient(context *gin.Context, manager *client.WebSocketManager) {
	socket, err := upgrader.Upgrade(context.Writer, context.Request, nil)

	if err != nil {
		log.Printf("createdWebsocket failed: %s", err.Error())
		return
	}

	manager.AddClient(socket)
}

func startCLI(manager *client.WebSocketManager) {
	routerCLI := gin.Default()

	routerCLI.GET("/cli", func(context *gin.Context) {
		websocketCLI(context, manager)
	})

	routerCLI.Run("127.0.0.1:81")
}

func setupViews(router *gin.Engine, manager *client.WebSocketManager) {

	router.GET("/", views.Index)

	router.GET("/websocket", func(context *gin.Context) {
		websocketClient(context, manager)
	})

	router.POST("/api/register", func(c *gin.Context) {
		api.Register(c, manager.DB)
	})

	router.Static("/game", "./game")
	router.LoadHTMLGlob("./templates/*")
}

func main() {
	mapData := gamemaps.NewMapData("./maps/test.json")
	fmt.Printf("mapData.Map: %v\n", mapData.Map)
	manager := client.NewWebSocketManager()
	manager.DB = database.SetupDB()
	router := gin.Default()

	go startCLI(manager)
	setupViews(router, manager)
	router.Run(":80")
}
