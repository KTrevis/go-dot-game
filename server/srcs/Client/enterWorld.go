package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"server/database"
	gamemaps "server/maps"
)

func (this *Client) enterWorld() error {
	if !this.authenticated {
		return errors.New("tried to enter world without authenticating")
	}

	var data struct {
		Character string
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		return fmt.Errorf("unmarshal failed: %s", this.body)
	}

	conn, _ := this.manager.DB.Acquire(context.TODO())
	defer conn.Release()

	this.character = database.GetCharacterByName(conn, data.Character, this.user.ID)

	if this.character == nil {
		const msg = "failed to find character"
		this.disconnect(msg)
		return errors.New(msg)
	}

	this.sendMessage("ENTER_WORLD", &Dictionary{
		"character": this.character,
		"map": gamemaps.NewMapData("test").Map,
	})
	return nil
}
