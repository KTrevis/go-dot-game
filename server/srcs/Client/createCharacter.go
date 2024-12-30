package client

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"server/classes"
	"server/database"
)


func (this *Client) createCharacter() error {
	if this.authenticated == false {
		const msg = "tried to create a character without authenticating"
		this.sendMessage(&Dictionary{"error": msg})
		return errors.New(msg)
	}

	var data struct {
		Class string
		Name string
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		this.disconnect()
		return errors.New("Client.createCharacter: unmarshal failed")
	}


	character := database.Character{
		UserID: this.user.ID,
		Name: data.Name,
		Level: 1,
		XP: 0,
		Class: classes.GetClass(data.Class),
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	if err = character.Create(conn); err != nil {
		this.sendMessage(&Dictionary{"error": err.Error()})
		return err
	}

	this.sendMessage(&Dictionary{"success": "character created"})
	log.Printf("%s character created", this.getFormattedIP())
	return nil
}
