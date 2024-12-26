package database

import (
	"context"
	"fmt"
	"log"
	"strings"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `database:"username"`
	Password string `database:"password"`
}

func hashString(password *string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Printf("hashString: failed")
		return
	}
	*password = string(bytes)
}

func (this *User) passwordIsValid(found *User) bool {
	err := bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(this.Password))
	return err == nil
}

func (this *User) usernameTaken(c *context.Context, db *DB) bool {
	res := db.QueryRow(*c, "SELECT username FROM users WHERE username=$1;", this.Username)
	var found string
	err := res.Scan(&found)
	return err == nil
}

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
	_, err = db.Exec(*c, "INSERT INTO users (username, password) VALUES ($1, $2);", this.Username, this.Password)

	return err
}

// Returns a user if the instance that used this method has valid username and password.
// Otherwise, returns nil and an error.
// func (user *User) login(db *mongo.Database) (*User, error) {
// 	var query = db.Collection("users").FindOne(context.TODO(), bson.M{"username": user.Username})
//
// 	if query.Err() != nil {
// 		return nil, fmt.Errorf("%s no user found with this username", user.Username)
// 	}
//
// 	var found User
// 	query.Decode(&found)
//
// 	if !user.passwordIsValid(&found) {
// 		return nil, fmt.Errorf("invalid password")
// 	}
//
// 	return user, nil
// }
