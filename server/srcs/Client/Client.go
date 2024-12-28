package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/classes"
	"server/database"
	"strings"
	"github.com/gorilla/websocket"
)

type Client struct {
	user			database.User
	msgType 		string
	socket  		*websocket.Conn
	manager 		*WebSocketManager
	authenticated	bool
	body			string
}

type Dictionary map[string]any

func (this *Client) getFormattedIP() string {
	str := fmt.Sprintf("[%s", this.socket.RemoteAddr())

	if this.user.Username != "" {
		str += fmt.Sprintf(" %s", this.user.Username)
	}

	str += "]"
	return str
}

func (this *Client) disconnect() {
	this.manager.RemoveClient(this.socket)
}

func (this *Client) treatMessage() {
	var err error
	const msg = "%s received message type %s"
	log.Printf(msg, this.getFormattedIP(), this.msgType)
	defer log.Printf("%s disconnected", this.getFormattedIP())

	switch this.msgType {
	case "LOGIN":
		err = this.login()

	case "GET_CLASSES":
		if this.authenticated == false {
			this.disconnect()
			return
		}
		this.sendMessage(&Dictionary{"classes": classes.GetClassesName()})

	case "CREATE_CHARACTER":

	default:
		const msg = "%s unknown message type %s, disconnecting client"
		log.Printf(msg, this.getFormattedIP(), this.msgType)
		this.disconnect()
		return
	}

	if err != nil {
		log.Printf("%s Client.treatMessage: %s", this.getFormattedIP(), err.Error())
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
	log.Printf("%s Client.Loop: new websocket connected", this.getFormattedIP())
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
