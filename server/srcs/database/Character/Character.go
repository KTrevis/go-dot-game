package character

import (
	"context"
	"fmt"
	"server/chunks"
	"server/classes"
	"server/classes/base"
	"server/database"
	"server/utils"
	"time"
)

type Character struct {
	UserID			int
	Name			string
	Level			int
	XP				int
	Position		utils.Vector2i
	Class			base_class.IBaseClass
	TilesPerSecond	int
	LastMovement	time.Time
}

func GetCharacterByName(db *database.DB, name string, userID int) *Character {
	var character Character
	class := ""

	const query = "SELECT (user_id, name, level, xp), (x, y), class FROM characters WHERE name=$1 AND user_id=$2;"
	err := db.QueryRow(context.TODO(), query, name, userID).Scan(
		&character, 
		&character.Position, &class,
	)

	if err != nil {
		fmt.Printf("GetByName: failed to find character %s", name)
		return nil
	}

	character.Class = classes.GetClass(class)
	character.TilesPerSecond = 5
	return &character
}

func (this *Character) GetChunk() *utils.Vector2i {
	return &utils.Vector2i{
		X: this.Position.X / (chunks.CHUNK_SIZE - 1),
		Y: this.Position.Y / (chunks.CHUNK_SIZE - 1),
	}
}

func (this *Character) IsOnChunkEdge(futurePos *utils.Vector2i) bool {
	if futurePos.X != this.Position.X &&
		futurePos.X % (chunks.CHUNK_SIZE / 2) == 0 {
		return true
	}

	if futurePos.Y != this.Position.Y &&
		futurePos.Y % (chunks.CHUNK_SIZE / 2) == 0 {
		return true
	}
	return false
}

func (this *Character) GetSurroundingChunks() *[]utils.Vector2i {
	chunk := this.GetChunk()

	return &[]utils.Vector2i{
		{X: chunk.X, Y: chunk.Y - 1},		// top
		{X: chunk.X + 1, Y: chunk.Y - 1},	// upper right
		{X: chunk.X + 1, Y: chunk.Y}, 	// right
		{X: chunk.X + 1, Y: chunk.Y + 1},	// bottom right
		{X: chunk.X, Y: chunk.Y + 1},		// bottom
		{X: chunk.X - 1, Y: chunk.Y + 1}, // bottom left
		{X: chunk.X - 1, Y: chunk.Y},		// left
		{X: chunk.X - 1, Y: chunk.Y - 1},	// upper left
	}
}
