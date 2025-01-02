package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"server/classes"
	"server/database"
)


func (this *Client) createCharacter() error {
	if this.authenticated == false {
		const msg = "tried to create a character without authenticating"
		this.sendMessage("CREATE_CHARACTER", &Dictionary{"error": msg})
		return errors.New(msg)
	}

	var data struct {
		Class string
		Name string
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		this.disconnect("invalid payload")
		return fmt.Errorf("unmarshal failed: %s", this.body)
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
		this.sendMessage("CREATE_CHARACTER", &Dictionary{"error": err.Error()})
		return err
	}

	this.sendMessage("CREATE_CHARACTER", &Dictionary{"success": "character created"})
	log.Printf("%s character created", this.getFormattedIP())
	return nil
}
