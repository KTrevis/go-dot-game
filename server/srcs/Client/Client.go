package client

import (
	"encoding/json"
	"log"
	"server/database"

	"github.com/gorilla/websocket"
)

type Client struct {
	user	*database.User
	msgType string
	socket  *websocket.Conn
	manager *WebSocketManager
	token	string
}

type Dictionary map[string]any


func (this *Client) Loop() {
	for {
		_, message, err := this.socket.ReadMessage()
		if err != nil {
			this.manager.RemoveClient(this.socket)
			log.Printf("client %s disconnected", this.user.Username)
			return
		}

		err = json.Unmarshal(message, &this.msgType)
		if err != nil {
			log.Printf("client.Loop failed to parse message: %s", message)
			continue
		}
		log.Printf("Received message of type %s from client %v", this.msgType, this.user)

		switch this.msgType {
		case "LOGIN":
			this.login()
		default:
			log.Printf("Received unknown message type, disconnecting client")
		}
		this.msgType = ""
	}
}

func (this *Client) sendMessage(msg *Dictionary) {
	str, _ := json.Marshal(msg)
	this.socket.WriteMessage(websocket.TextMessage, str)
}

func (this *Client) login() {
	_, message, err := this.socket.ReadMessage()

	if err != nil {
		log.Printf("client disconnected %v", this.user.Username)
		this.manager.RemoveClient(this.socket)
		return
	}

	if this.token != "" {
		this.sendMessage(&Dictionary{"error": "you are already logged in"})
		return
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		log.Printf("client.login failed to parse message: %s", message)
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	this.token, err = credentials.Login(this.manager.DB)

	if err != nil {
		log.Printf("client.login failed: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	this.user = &credentials
	log.Printf("client.log success: %s", this.user.Username)
	this.sendMessage(&Dictionary{"token": this.token})
	this.manager.AddOnlineUser(this.user)
}
