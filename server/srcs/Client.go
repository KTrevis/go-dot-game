package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

type Dictionary map[string]interface{}
type MessageType int

const (
	NONE MessageType = iota
	REGISTER_LOGIN
)

func MessageTypeToString(msgType MessageType) string {
	switch msgType {
	case MessageType(NONE):
		return "NONE"
	case MessageType(REGISTER_LOGIN):
		return "REGISTER_LOGIN"
	}
	return "UNKNOWN"
}

type Client struct {
	user    *User
	msgType MessageType
	socket  *websocket.Conn
	manager *WebSocketManager
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
			log.Printf("Received message of type %s from client %s", MessageTypeToString(client.msgType), client.user)
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

func (client *Client) sendMessage(msg *Dictionary) {
	str, _ := json.Marshal(msg)
	client.socket.WriteMessage(websocket.TextMessage, str)
}

func (client *Client) registerLogin() bool {
	client.msgType = MessageType(NONE)
	_, message, err := client.socket.ReadMessage()
	if err != nil {
		return false
	}

	var credentials User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		log.Printf("registerLogin: %s", err.Error())
		return false
	}

	id, err := credentials.RegisterLogin(client.manager.DB, client)

	if err != nil {
		client.sendMessage(&Dictionary{"error": err.Error()})
		return true
	}
	client.sendMessage(&Dictionary{"token": id})
	return true
}
