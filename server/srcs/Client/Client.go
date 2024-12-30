package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/database"
	"strings"
	"github.com/gorilla/websocket"
)

type Client struct {
	socket  		*websocket.Conn
	manager 		*WebSocketManager
	msgType 		string
	body			string
	authenticated	bool
	user			database.User
}

type Dictionary map[string]any

func (this *Client) getFormattedIP() string {
	str := fmt.Sprintf("[%s", this.socket.RemoteAddr())

	if this.authenticated {
		str += fmt.Sprintf(" %s", this.user.Username)
	}

	str += "]"
	return str
}

func (this *Client) disconnect() {
	this.manager.RemoveClient(this.socket)
}

func (this *Client) treatMessage() {
	const msg = "%s received message type %s"
	log.Printf(msg, this.getFormattedIP(), this.msgType)

	m := map[string]func()error {
		"LOGIN": this.login,
		"GET_CLASS_LIST": this.getClassList,
		"CREATE_CHARACTER": this.createCharacter,
		"DELETE_CHARACTER": this.deleteCharacter,
		"GET_CHARACTER_LIST": this.getCharacterList,
	}

	fn, ok := m[this.msgType]
	
	if !ok {
		const msg = "%s unknown message type %s, disconnecting client"
		log.Printf(msg, this.getFormattedIP(), this.msgType)
		this.disconnect()
		return
	}

	err := fn()

	if err != nil {
		const msg = "%s Client.treatMessage: %s"
		log.Printf(msg, this.getFormattedIP(), err.Error())
	}
}

func (this *Client) setMessageType(message []byte) error {
	split := strings.Split(string(message), "\r\n")

	if len(split) != 2 {
		const msg = "%s Client.setMessage missing header in request"
		log.Printf(msg, this.getFormattedIP())
		this.sendMessage(&Dictionary{"error": "invalid request"})
		return errors.New(msg)
	}

	this.msgType = split[0]
	this.body = split[1]
	return nil
}

func (this *Client) Loop() {
	log.Printf("%s connected", this.getFormattedIP())
	defer log.Printf("%s disconnected", this.getFormattedIP())
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
