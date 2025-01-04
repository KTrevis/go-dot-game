package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"server/utils"
	"time"
)

func (this *Client) teleported(newPos *utils.Vector2i) bool {
	vec := this.character.Position.Sub(newPos)

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
	timeSinceMov += time.Millisecond * 100
	this.character.LastMovement = now
	timePerTile := time.Second / time.Duration(this.character.TilesPerSecond)

	return timeSinceMov < timePerTile
}

func (this *Client) sendSurroundingChunks() error {
	if this.character == nil {
		return errors.New("getting chunks without being ingame")
	}

	chunks := this.character.GetSurroundingChunks()

	for _, v := range *chunks {
		chunk, ok := this.chunks.Chunks[v]

		if ok {
			this.sendMessage("SEND_MAP", &Dict{
				"map": chunk,
			})
		}
	}
	return nil
}

func (this *Client) updatePlayerPosition() error {
	if this.character == nil {
		const msg = "tried changing position without being ingame"
		this.disconnect(msg)
		return errors.New("tried changing position without being ingame") 
	}

	var data struct {
		Position utils.Vector2i
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

	if this.character.IsOnChunkEdge(&data.Position) {
		this.sendSurroundingChunks()
	}

	this.character.Position = data.Position
	this.sendPosition(this.character)
	return nil
}
