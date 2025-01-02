package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/utils"
	"time"
)

func (this *Client) canUpdatePos() error {
	if this.character == nil {
		const msg = "tried changing position without being in game"
		this.disconnect(msg)
		return errors.New("tried changing position without being in game") 
	}

	if !this.authenticated {
		// this should never happen since the player
		// has to be authenticated to have a non nil character
		const msg = "tried changing position unauthenticated"
		this.disconnect(msg)
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

func (this *Client) movedTooFast() bool {
	now := time.Now()

	if this.character.LastMovement.IsZero() {
		this.character.LastMovement = now
		return false
	}

	timeSinceMov := now.Sub(this.character.LastMovement)
	timeSinceMov += time.Millisecond * 30
	this.character.LastMovement = now
	timePerTile := time.Second / time.Duration(this.character.TilesPerSecond)

	return timeSinceMov < timePerTile
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
		const msg = "tried to teleport"
		this.disconnect(msg)
		return fmt.Errorf(msg)
	}

	if this.movedTooFast() {
		const msg = "moved too fast"
		this.disconnect(msg)
		return errors.New(msg)
	}

	this.character.Position = data.Position

	return nil
}
