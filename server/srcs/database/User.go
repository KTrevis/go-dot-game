package database

import (
	"context"
	"fmt"
	"log"
	"server/utils"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username 	string `database:"username"`
	Password 	string `database:"password"`
}

func hashString(password *string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Printf("hashString: failed")
		return
	}
	*password = string(bytes)
}

// Returns nil if user password is valid.
func (this *User) passwordIsValid(db *DB) error {
	var hash string
	err := db.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1;", this.Username).Scan(&hash)
	
	if err != nil {
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(this.Password)) != nil {
		return fmt.Errorf("invalid password")
	}

	return nil
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

// Returns a random string to use as a session token if the user successfully logged in.
// Returns an empty string and an error otherwise.
func (this *User) Login(db *DB) (string, error) {
	var hash string
	err := db.QueryRow(context.Background(), "SELECT password FROM users WHERE username=$1;", this.Username).Scan(&hash)
	
	if err != nil {
		return "", fmt.Errorf("user %s not found", this.Username)
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(this.Password)) != nil {
		return "", fmt.Errorf("invalid password")
	}
	this.Password = hash
	return utils.RandStr(), nil
}
