package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password string
}

func hashString(password *string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Printf("hashString: failed")
		return
	}
	*password = string(bytes)
}

func (user *User) passwordIsValid(found *User) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password))
	return err == nil
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
