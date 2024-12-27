package database

import (
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username 	string	`db:"username"`
	Password 	string 	`db:"password"`
	ID			int		`db:"id"`
}

func hashString(password *string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Printf("hashString: failed")
		return
	}
	*password = string(bytes)
}

func (this *User) usernameTaken(c *context.Context, db *DB) bool {
	var found string

	res := db.QueryRow(*c, "SELECT username FROM users WHERE username=$1;", this.Username)
	err := res.Scan(&found)
	return err == nil
}

// Returns nil if the user has been successfully added to the database.
func (this *User) AddToDB(c *context.Context, db *DB) error {
	const MINIMUM_LEN = 4

	if len(this.Username) < MINIMUM_LEN {
		return fmt.Errorf("username must be at least %d characters long", MINIMUM_LEN)
	}

	if len(this.Password) < MINIMUM_LEN {
		return fmt.Errorf("password must be at least %d characters long", MINIMUM_LEN)
	}

	if strings.Index(this.Username, " ") != -1 {
		return fmt.Errorf("spaces are not allowed in username")
	}

	if this.usernameTaken(c, db) {
		return fmt.Errorf("username already taken")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(this.Password), 14)

	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	this.Password = string(hash)
	const query = "INSERT INTO users (username, password) VALUES ($1, $2);"
	_, err = db.Exec(*c, query, this.Username, this.Password)

	return err
}

func (this *User) Login(db *DB) error {
	var user User
	err := db.QueryRow(context.Background(), "SELECT (username, password, id) FROM users WHERE username=$1;", this.Username).Scan(&user)
	
	if err != nil {
		return fmt.Errorf("user %s not found", this.Username)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(this.Password)) != nil {
		return fmt.Errorf("invalid password")
	}
	*this = user
	return nil
}
