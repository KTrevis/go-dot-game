package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/database"
)

func (this *Client) login() error {
	_, message, err := this.socket.ReadMessage()

	if err != nil {
		log.Printf("Client.login: %s", err.Error())
		return err
	}

	if  this.authenticated {
		msg := fmt.Sprintf("%s client %s tried to log in while already authenticated", this.socket.RemoteAddr(), this.user.Username)
		this.sendMessage(&Dictionary{"error": "you are already authenticated"})
		return errors.New(msg)
	}

	var credentials database.User
	err = json.Unmarshal(message, &credentials)

	if err != nil {
		msg := fmt.Sprintf("client.login failed to unmarshal message: %s", message)
		this.sendMessage(&Dictionary{"error": err.Error()})
		this.disconnect()
		return errors.New(msg)
	}

	err = credentials.Login(this.manager.DB)

	if err != nil {
		msg := fmt.Sprintf("credentials.Login: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return errors.New(msg)
	}

	if this.manager.UserIsOnline(&credentials) {
		this.sendMessage(&Dictionary{"error": "this account is already logged in"})
		const msg = "credentials.Login: tried to log in to already active session %v %v"
		return fmt.Errorf(msg, credentials.Username, this.socket.RemoteAddr())
	}

	this.user = credentials
	this.authenticated = true
	this.sendMessage(&Dictionary{"authenticated": true})
	this.manager.AddOnlineUser(&this.user)
	log.Printf("Client authenticated: %s", this.user.Username)
	return nil
}
