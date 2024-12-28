package client

import (
	"encoding/json"
	"log"
	"server/database"

	"github.com/gorilla/websocket"
)

type Client struct {
	user			database.User
	msgType 		string
	socket  		*websocket.Conn
	manager 		*WebSocketManager
	authenticated	bool
}

type Dictionary map[string]any

func (this *Client) disconnect() {
	this.manager.RemoveClient(this.socket)
}

func (this *Client) treatMessage() {
	var err error
	const msg = "%s Received message type %s from client %s"
	log.Printf(msg, this.socket.RemoteAddr(), this.msgType, this.user.Username)

	switch this.msgType {
	case "LOGIN":
		err = this.login()

	case "GET_CLASSES":

	default:
		const msg = "%s Unknown message type %s, disconnecting client %s"
		log.Printf(msg, this.socket.RemoteAddr(), this.msgType, this.user.Username)
		this.disconnect()
		return
	}

	if err != nil {
		log.Printf("Client.treatMessage: %s", err.Error())
	}
}

func (this *Client) setMessageType(message []byte) error {
	err := json.Unmarshal(message, &this.msgType)

	if err != nil {
		const msg = "client.setMessage failed to unmarshal message: %s"
		log.Printf(msg, message)
		this.disconnect()
		return err
	}
	return nil
}

func (this *Client) Loop() {
	log.Printf("Client.Loop: new websocket connected %v", this.socket.RemoteAddr())
	for {
		_, message, err := this.socket.ReadMessage()

		if err != nil {
			this.disconnect()
			return
		}

		if this.setMessageType(message) != nil {
			this.disconnect()
			return
		}

		this.treatMessage()
	}
}

func (this *Client) sendMessage(msg *Dictionary) {
	str, _ := json.Marshal(msg)
	this.socket.WriteMessage(websocket.TextMessage, str)
}
