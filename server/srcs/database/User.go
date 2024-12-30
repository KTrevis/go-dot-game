package database

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username 	string	`db:"username"`
	Password 	string 	`db:"password"`
	ID			int		`db:"id"`
}

func (this *User) isValid() error {
	const MINIMUM_LEN = 4

	if len(this.Username) < MINIMUM_LEN {
		const format = "username must be at least %d characters long"
		return fmt.Errorf(format, MINIMUM_LEN)
	}

	if len(this.Password) < MINIMUM_LEN {
		const format = "password must be at least %d characters long"
		return fmt.Errorf(format, MINIMUM_LEN)
	}

	if strings.Index(this.Username, " ") != -1 {
		const format = "spaces are not allowed in username"
		return fmt.Errorf(format)
	}

	return nil
}

func (this *User) CreateAccount(db *DB) error {
	if err := this.isValid(); err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(this.Password), 10)

	if err != nil {
		return fmt.Errorf("failed to hash password")
	}

	this.Password = string(hash)
	const query = "INSERT INTO users (username, password) VALUES ($1, $2);"
	_, err = db.Exec(context.TODO(), query, this.Username, this.Password)

	if err != nil {
		return errors.New("username already taken")
	}

	return nil
}

func (this *User) Login(db *DB, onlineUsers map[int]bool) error {
	var user User
	const query = "SELECT (username, password, id) FROM users WHERE username=$1;"
	err := db.QueryRow(context.TODO(), query, this.Username).Scan(&user)
	
	if err != nil {
		return fmt.Errorf("account %s not found", this.Username)
	}

	if _, ok := onlineUsers[user.ID]; ok {
		return errors.New("this account is already logged in")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(this.Password)) != nil {
		return fmt.Errorf("invalid password")
	}
	*this = user
	return nil
}
