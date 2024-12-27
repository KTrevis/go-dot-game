package client

import (
	"encoding/json"
	"fmt"
	"log"
	"server/database"
)

func (this *Client) login() error {
	_, message, err := this.socket.ReadMessage()

	if err != nil {
		log.Println("Client.login: %s", err.Error())
		return err
	}

	if  this.authenticated {
		msg := fmt.Sprintf("client %s tried to log in while already authenticated", this.user.Username)
		log.Println(msg)
		this.sendMessage(&Dictionary{"error": "you are already authenticated"})
		return fmt.Errorf(msg)
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		msg := fmt.Sprintf("client.login failed to unmarshal message: %s", message)
		log.Println(msg)
		this.sendMessage(&Dictionary{"error": err.Error()})
		return fmt.Errorf(msg)
	}

	err = credentials.Login(this.manager.DB)

	if err != nil {
		msg := fmt.Sprintf("credentials.Login failed: %s", err.Error())
		log.Println(msg)
		this.sendMessage(&Dictionary{"error": err.Error()})
		return fmt.Errorf(msg)
	}

	this.user = credentials
	this.authenticated = true
	this.sendMessage(&Dictionary{"authenticated": true})
	this.manager.AddOnlineUser(&this.user)
	log.Printf("Client authenticated: %s", this.user.Username)
	return nil
}
