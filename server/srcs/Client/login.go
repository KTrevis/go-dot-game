package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/database"
)

func (this *Client) login() error {
	if  this.authenticated {
		msg := fmt.Sprintf("client tried to log in while already authenticated")
		this.sendMessage(&Dictionary{"error": "you are already authenticated"})
		return errors.New(msg)
	}

	var credentials database.User
	err := json.Unmarshal([]byte(this.body), &credentials)

	if err != nil {
		msg := fmt.Sprintf("client.login failed to unmarshal message: %s", this.body)
		this.sendMessage(&Dictionary{"error": err.Error()})
		this.disconnect()
		return errors.New(msg)
	}

	// if this.manager.UserIsOnline(&credentials) {
	// 	this.sendMessage(&Dictionary{"error": "this account is already logged in"})
	// 	const msg = "credentials.Login: tried to login to already active session %s"
	// 	return fmt.Errorf(msg, credentials.Username)
	// }

	err = credentials.Login(this.manager.DB, this.manager.onlineUsers)

	if err != nil {
		msg := fmt.Sprintf("credentials.Login: %s", err.Error())
		this.sendMessage(&Dictionary{"error": err.Error()})
		return errors.New(msg)
	}

	this.user = credentials
	this.authenticated = true
	this.sendMessage(&Dictionary{"authenticated": true})
	this.manager.AddOnlineUser(&this.user)
	log.Printf("%s authenticated", this.getFormattedIP())
	return nil
}
