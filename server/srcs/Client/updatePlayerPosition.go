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

func (this *Client) sendSingleChunk(pos *utils.Vector2i) {
	chunk := this.chunks.Chunks[*pos]
	if chunk != nil {
		this.sendMessage("SEND_MAP", &Dict{
			"map": chunk,
		})
	}
}

func (this *Client) sendNearChunks() {
	chunkPos := this.character.ConvertPosToChunk()
	directions := []utils.Vector2i{
		{X: chunkPos.X, Y: chunkPos.Y - 1},		// top
		{X: chunkPos.X + 1, Y: chunkPos.Y - 1},	// upper right
		{X: chunkPos.X + 1, Y: chunkPos.Y}, 	// right
		{X: chunkPos.X + 1, Y: chunkPos.Y + 1},	// bottom right
		{X: chunkPos.X, Y: chunkPos.Y + 1},		// bottom
		{X: chunkPos.X - 1, Y: chunkPos.Y + 1}, // bottom left
		{X: chunkPos.X - 1, Y: chunkPos.Y},		// left
		{X: chunkPos.X - 1, Y: chunkPos.Y - 1},	// upper left
	}

	for _, v := range directions {
		this.sendSingleChunk(&v)
	}
}

func (this *Client) updatePlayerPosition() error {
	if err := this.canUpdatePos(); err != nil {
		return err
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
		this.sendNearChunks()
	}

	this.character.Position = data.Position

	return nil
}
