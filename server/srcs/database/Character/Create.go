package character

import (
	"context"
	"errors"
	"server/database"
	"strings"
)

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

func (this *Character) Create(db *database.DB) error {
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
