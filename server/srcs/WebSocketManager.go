package main

import (
	"sync"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

type WebSocketManager struct {
	Clients map[*websocket.Conn]*Client
	DB      *mongo.Database
	mutex   sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Clients: make(map[*websocket.Conn]*Client),
	}
}

func (manager *WebSocketManager) AddClient(socket *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	manager.Clients[socket] = &Client{
		manager: manager,
		socket:  socket,
	}
}

func (manager *WebSocketManager) RemoveClient(socket *websocket.Conn) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	delete(manager.Clients, socket)
	socket.Close()
}
