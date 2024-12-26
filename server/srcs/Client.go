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
	LOGIN
)

func MessageTypeToString(msgType MessageType) string {
	switch msgType {
	case MessageType(NONE):
		return "NONE"
	case MessageType(LOGIN):
		return "LOGIN"
	}
	return "UNKNOWN"
}

type Client struct {
	user    *User
	msgType MessageType
	socket  *websocket.Conn
	logged	bool
	manager *WebSocketManager
}

func (this *Client) Loop() {
	for {
		_, message, err := this.socket.ReadMessage()
		if err != nil {
			log.Printf("client.Loop: failed to read message")
			continue
		}

		err = json.Unmarshal(message, &this.msgType)
		if err != nil {
			log.Printf("client.Loop: failed to parse message")
			continue
		}
		log.Printf("Received message of type %s from client %s", MessageTypeToString(this.msgType), this.user)

		switch this.msgType {
		case LOGIN:
			this.login()
		}
		this.msgType = NONE
	}
}

func (this *Client) sendMessage(msg *Dictionary) {
	str, _ := json.Marshal(msg)
	this.socket.WriteMessage(websocket.TextMessage, str)
}

func (this *Client) login() {
	_, message, err := this.socket.ReadMessage()

	if err != nil {
		log.Printf("client.login failed to read message: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	var credentials User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		log.Printf("client.login failed to parse json: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	// TODO: check login here
	if err != nil {
		log.Printf("client.login failed to login: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}
	this.logged = true
}
