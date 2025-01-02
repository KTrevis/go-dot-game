package character

import (
	"context"
	"fmt"
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

func (this *Character) ConvertPosToChunk() *utils.Vector2i {
	const CHUNK_SIZE = 50

	return &utils.Vector2i{
		X: this.Position.X / CHUNK_SIZE - 1,
		Y: this.Position.Y / CHUNK_SIZE - 1,
	}
}
