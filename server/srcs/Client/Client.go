package client

import (
	"encoding/json"
	"fmt"
	"log"
	"server/database"

	"github.com/gorilla/websocket"
)

type Client struct {
	user	database.User
	msgType string
	socket  *websocket.Conn
	manager *WebSocketManager
	token	string
}

type Dictionary map[string]any

func (this *Client) disconnect() {
	this.manager.RemoveClient(this.socket)
}

func (this *Client) treatMessage() error {
	switch this.msgType {
	case "LOGIN":
		this.login()

	case "GET_CLASSES":

	default:
		const msg = "Unknown message type, disconnecting client %s"
		this.disconnect()
		return fmt.Errorf("unknown message type")
	}
	const msg = "Received message of type %s from client %s"
	log.Printf(msg, this.msgType, this.user.Username)
	return nil
}

func (this *Client) Loop() {
	log.Printf("Client.Loop: new websocket connected")
	for {
		_, message, err := this.socket.ReadMessage()
		if err != nil {
			this.disconnect()
			return
		}

		err = json.Unmarshal(message, &this.msgType)

		if err != nil {
			const msg = "client.Loop failed to unmarshal message: %s"
			log.Printf(msg, message)
			this.disconnect()
			return
		}
		
		if this.treatMessage() != nil {
			return
		}
	}
}

func (this *Client) sendMessage(msg *Dictionary) {
	str, _ := json.Marshal(msg)
	this.socket.WriteMessage(websocket.TextMessage, str)
}
