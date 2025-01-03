package client

import (
	"server/chunks"
	"server/database"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Clients map[*websocket.Conn]*Client
	onlineUsers map[int]bool
	DB		*database.DBPool
	mutex   sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		DB: database.SetupDB(),
		Clients: make(map[*websocket.Conn]*Client),
		onlineUsers: make(map[int]bool),
	}
}

func (this *WebSocketManager) AddClient(socket *websocket.Conn, chunks *chunks.ChunkHandler) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.Clients[socket] = &Client{
		manager: this,
		socket:  socket,
		authenticated: false,
		chunks: chunks,
	}
	go this.Clients[socket].Loop()
}

func (this *WebSocketManager) removeOnlineUser(socket *websocket.Conn) {
	user := &this.Clients[socket].user
	_, ok := this.onlineUsers[user.ID]

	if ok {
		delete(this.onlineUsers, user.ID)
	}
}

func (this *WebSocketManager) RemoveClient(socket *websocket.Conn, reason string) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	_, ok := this.Clients[socket]

	if !ok {
		return
	}

	this.removeOnlineUser(socket)
	delete(this.Clients, socket)
	socket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(1000, reason))
	socket.Close()
}

func (this *WebSocketManager) UserIsOnline(user *database.User) bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	_, ok := this.onlineUsers[user.ID]
	return ok
}

func (this *WebSocketManager) AddOnlineUser(user *database.User) {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	this.onlineUsers[user.ID] = true
}
