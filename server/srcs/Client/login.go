package client

import (
	"encoding/json"
	"log"
	"server/database"
)

func (this *Client) login() {
	_, message, err := this.socket.ReadMessage()

	if err != nil {
		log.Printf("client disconnected: %v", this.user.Username)
		this.disconnect()
		return
	}

	if this.token != "" {
		const msg = "client %s tried to log in while already logged in"
		log.Printf(msg, this.user.Username)
		this.sendMessage(&Dictionary{"error": "you are already logged in"})
		return
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		log.Printf("client.login failed to unmarshal message: %s", message)
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	this.token, err = credentials.Login(this.manager.DB)

	if err != nil {
		log.Printf("credentials.Login failed: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	this.user = credentials
	log.Printf("Client authenticated: %s", this.user.Username)
	this.sendMessage(&Dictionary{"token": this.token})
	this.manager.AddOnlineUser(&this.user)
}
