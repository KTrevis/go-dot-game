package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/utils"
)

func (this *Client) canUpdatePos() error {
	if this.character == nil {
		this.disconnect()
		return errors.New("tried changing position without being in game") 
	}

	if !this.authenticated {
		// this should never happen since the player
		// has to be authenticated to have a non nil character
		this.disconnect()
		const msg = "not authenticated"
		return errors.New(msg) 
	}
	return nil
}

func (this *Client) teleported(newPos *utils.Vector2) bool {
	vec := this.character.Position.Substract(newPos)

	if vec.X < -1 || vec.X > 1 {
		return true
	}

	if vec.Y < -1 || vec.Y > 1 {
		return true
	}

	return false
}

func (this *Client) updatePlayerPosition() error {
	if err := this.canUpdatePos(); err != nil {
		return err
	}

	var data struct {
		Position utils.Vector2
	}

	err := json.Unmarshal([]byte(this.body), &data)

	if err != nil {
		return fmt.Errorf("unmarshal failed: %s", this.body)
	}

	if this.teleported(&data.Position) {
		this.disconnect()
		return fmt.Errorf("tried to teleport")
	}

	this.character.Position = data.Position
	return nil
}
