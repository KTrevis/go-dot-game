package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"server/database/Character"
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

	this.character = character.GetCharacterByName(conn, data.Character, this.user.ID)

	if this.character == nil {
		const msg = "failed to find character"
		this.disconnect(msg)
		return errors.New(msg)
	}
	fmt.Printf("this.character.ConvertPosToChunk(): %v\n", this.character.ConvertPosToChunk())
	gamemap := this.chunks.Chunks[*this.character.ConvertPosToChunk()]
	if gamemap == nil {
		fmt.Printf("gamemap: %v\n", gamemap)
		return nil
	}
	this.sendMessage("ENTER_WORLD", &Dict{
		"character": this.character,
		"map": gamemap,
	})
	return nil
}
