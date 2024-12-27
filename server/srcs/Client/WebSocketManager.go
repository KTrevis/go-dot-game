package client

import (
	"log"
	"server/database"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Clients map[*websocket.Conn]*Client
	OnlineUsers map[*database.User]bool
	DB		*database.DB
	mutex   sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Clients: make(map[*websocket.Conn]*Client),
		OnlineUsers: make(map[*database.User]bool),
	}
}

func (this *WebSocketManager) AddClient(socket *websocket.Conn) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.Clients[socket] = &Client{
		manager: this,
		socket:  socket,
	}
}

func (this *WebSocketManager) removeOnlineUser(socket *websocket.Conn) {
	user := &this.Clients[socket].user
	_, ok := this.OnlineUsers[user]

	if ok {
		delete(this.OnlineUsers, user)
	}
}

func (this *WebSocketManager) RemoveClient(socket *websocket.Conn) {
	this.mutex.Lock()
	client := this.Clients[socket]

	log.Printf("client %s disconnected", client.user.Username)

	this.removeOnlineUser(socket)
	delete(this.Clients, socket)
	socket.Close()
	this.mutex.Unlock()
}

func (this *WebSocketManager) AddOnlineUser(user *database.User) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.OnlineUsers[user] = true
}
