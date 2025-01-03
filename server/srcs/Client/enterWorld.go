package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"server/chunks"
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

	gamemap := this.chunks.Chunks[*this.character.ConvertPosToChunk()]

	if gamemap == nil {
		this.character.Position = chunks.SPAWN
		gamemap = this.chunks.Chunks[this.character.Position]
	}

	this.sendMessage("ENTER_WORLD", &Dict{
		"character": this.character,
		"map": gamemap,
	})
	return nil
}
