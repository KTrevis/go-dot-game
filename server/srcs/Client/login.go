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

	if  this.authenticated {
		const msg = "client %s tried to log in while already authenticated"
		log.Printf(msg, this.user.Username)
		this.sendMessage(&Dictionary{"error": "you are already authenticated"})
		return
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		log.Printf("client.login failed to unmarshal message: %s", message)
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	err = credentials.Login(this.manager.DB)

	if err != nil {
		log.Printf("credentials.Login failed: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return
	}

	this.user = credentials
	this.authenticated = true
	log.Printf("Client authenticated: %s", this.user.Username)
	this.sendMessage(&Dictionary{"authenticated": true})
	this.manager.AddOnlineUser(&this.user)
}
