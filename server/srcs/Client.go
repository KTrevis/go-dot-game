package srcs

import (
	"encoding/json"
	"log"
	"server/srcs/database"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageType int

const (
	NONE MessageType = iota
	REGISTER_LOGIN
)

type Client struct {
	username    string
	msgType     MessageType
	socket      *websocket.Conn
	manager     *WebSocketManager
	mongoClient *mongo.Client
}

func (client *Client) Loop() {
	for {
		_, message, err := client.socket.ReadMessage()
		if err != nil {
			client.manager.RemoveClient(client.socket)
			return
		}

		if client.msgType == MessageType(NONE) {
			err = json.Unmarshal(message, &client.msgType)
			if err != nil {
				continue
			}
			log.Printf("received message of type %d", MessageType(client.msgType))
		}

		switch client.msgType {
		case REGISTER_LOGIN:
			if !client.registerLogin() {
				client.manager.RemoveClient(client.socket)
				return
			}
		}
	}
}

func (client *Client) registerLogin() bool {
	client.msgType = MessageType(NONE)
	_, message, err := client.socket.ReadMessage()
	if err != nil {
		return false
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)
	if err != nil {
		return false
	}

	if credentials.UsernameTaken(client.manager.DB) {
		return true
	}
	credentials.CreateUser(client.manager.DB)
	return true
}
