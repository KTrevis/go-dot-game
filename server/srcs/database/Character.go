package database

import (
	"context"
	"errors"
	"fmt"
	"server/classes"
	"server/classes/base"
	"server/utils"
	"strings"
)

type Character struct {
	UserID		int
	Name		string
	Level		int
	XP			int
	Position	utils.Vector2
	Class		base_class.IBaseClass
}

func (this *Character) isValid() error {
	if this.Class == nil {
		return errors.New("tried to create invalid class")
	}

	if strings.Index(this.Name, " ") != -1 {
		return errors.New("invalid character name")
	}

	if len(this.Name) < 4 {
		return errors.New("character name must be at least 4 characters long")
	}
	return nil
}

func (this *Character) Create(db *DB) error {
	err := this.isValid()

	if err != nil {
		return err
	}

	query := "INSERT INTO characters (user_id, name, level, xp, class)"
	query += " VALUES ($1, $2, $3, $4, $5);"
	_, err = db.Exec(context.TODO(), query,
	this.UserID, this.Name, this.Level, this.XP, this.Class.GetName())

	if err != nil {
		return errors.New("character name already taken")
	}

	return nil
}

func GetCharacterByName(db *DB, name string, userID int) *Character {
	var character Character
	class := ""

	const query = "SELECT (user_id, name, level, xp), x, y, class FROM characters WHERE name=$1 AND user_id=$2;"
	err := db.QueryRow(context.TODO(), query, name, userID).Scan(
		&character, 
		&character.Position.X, &character.Position.Y,
		&class,
	)

	if err != nil {
		fmt.Printf("GetByName: failed to find character %s", name)
		return nil
	}

	character.Class = classes.GetClass(class)
	return &character
}
