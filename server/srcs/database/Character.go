package database

import (
	"context"
	"errors"
	"server/classes/base"
	"strings"
)

type Character struct {
	UserID		int
	Name		string
	Level		int
	XP			int
	Position	[2]int
	Class		base_class.IBaseClass
}

func (this *Character) isValid() error {
	if this.Class == nil {
		return errors.New("tried to create invalid class")
	}

	if strings.Index(this.Name, " ") != -1 {
		return errors.New("invalid character name")
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
