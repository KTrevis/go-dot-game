package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/database"
)

func (this *Client) login() error {
	if  this.authenticated {
		msg := fmt.Sprintf("client tried to log in while already authenticated")
		this.sendMessage("LOGIN", &Dictionary{"error": "you are already authenticated"})
		return errors.New(msg)
	}

	var credentials database.User
	err := json.Unmarshal([]byte(this.body), &credentials)

	if err != nil {
		msg := fmt.Sprintf("unmarshal failed: %s", this.body)
		this.sendMessage("LOGIN", &Dictionary{"error": err.Error()})
		this.disconnect()
		return errors.New(msg)
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	err = credentials.Login(conn, this.manager.onlineUsers)

	if err != nil {
		msg := fmt.Sprintln(err.Error())
		this.sendMessage("LOGIN", &Dictionary{"error": err.Error()})
		return errors.New(msg)
	}


	this.user = credentials
	this.authenticated = true
	this.sendMessage("LOGIN", &Dictionary{"authenticated": true})
	this.manager.AddOnlineUser(&this.user)
	log.Printf("%s authenticated", this.getFormattedIP())
	return nil
}
