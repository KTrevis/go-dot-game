package client

import (
	"log"
	"server/database"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Clients map[*websocket.Conn]*Client
	onlineUsers map[int]bool
	DB		*database.DB
	mutex   sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		Clients: make(map[*websocket.Conn]*Client),
		onlineUsers: make(map[int]bool),
	}
}

func (this *WebSocketManager) AddClient(socket *websocket.Conn) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.Clients[socket] = &Client{
		manager: this,
		socket:  socket,
		authenticated: false,
	}
}

func (this *WebSocketManager) removeOnlineUser(socket *websocket.Conn) {
	user := &this.Clients[socket].user
	_, ok := this.onlineUsers[user.ID]

	if ok {
		delete(this.onlineUsers, user.ID)
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

func (this *WebSocketManager) UserIsOnline(user *database.User) bool {
	_, ok := this.onlineUsers[user.ID]
	return ok
}

func (this *WebSocketManager) AddOnlineUser(user *database.User) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.onlineUsers[user.ID] = true
}
