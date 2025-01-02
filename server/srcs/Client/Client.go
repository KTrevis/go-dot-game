package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/database"
	character "server/database/Character"
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
	character		*character.Character
}

type Dict map[string]any

func (this *Client) getFormattedIP() string {
	str := fmt.Sprintf("[%s", this.socket.RemoteAddr())

	if this.authenticated {
		str += fmt.Sprintf(" %s", this.user.Username)
	}

	str += "]"
	return str
}

func (this *Client) save() {
	if this.character == nil {
		return
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	const query =  "UPDATE characters SET level=$1, xp=$2, x=$3, y=$4 WHERE name=$5;"
	conn.Exec(context.TODO(), query,
		this.character.Level, this.character.XP,
		this.character.Position.X, this.character.Position.Y,
		this.character.Name)
}

func (this *Client) disconnect(reason string) {
	this.save()
	this.manager.RemoveClient(this.socket, reason)
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
		"GET_MAP": this.getMap,
		"ENTER_WORLD": this.enterWorld,
		"UPDATE_PLAYER_POSITION": this.updatePlayerPosition,
	}

	fn, ok := m[this.msgType]
	
	if !ok {
		const msg = "%s unknown message type %s, disconnecting client"
		log.Printf(msg, this.getFormattedIP(), this.msgType)
		this.disconnect("invalid payload")
		return
	}

	err := fn()

	if err != nil {
		const msg = "%s %s %s"
		log.Printf(msg, this.getFormattedIP(), this.msgType, err.Error())
	}
}

func (this *Client) setMessageType(message []byte) error {
	split := strings.Split(string(message), "\r\n")

	if len(split) != 2 {
		const msg = "%s Client.setMessage missing header in request"
		log.Printf(msg, this.getFormattedIP())
		return errors.New(msg)
	}

	this.msgType = split[0]
	this.body = split[1]
	return nil
}

func (this *Client) Loop() {
	log.Printf("%s connected", this.getFormattedIP())
	defer func() { 
		// we use an anonymous function because the string
		// was formatted before the function was called
		log.Printf("%s disconnected", this.getFormattedIP())
	}()

	for {
		_, message, err := this.socket.ReadMessage()

		if err != nil {
			this.disconnect("failed to read message")
			return
		}

		if this.setMessageType(message) != nil {
			this.disconnect("invalid payload")
			return
		}

		this.treatMessage()
	}
}

func (this *Client) sendMessage(msgType string, msg *Dict) {
	str, _ := json.Marshal(msg)
	msgType += fmt.Sprintf("\r\n%s", str)
	this.socket.WriteMessage(websocket.TextMessage, []byte(msgType))
}
