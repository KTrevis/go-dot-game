package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func hashPassword(password *string) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Printf("failed to hash password")
		return
	}
	*password = string(bytes)
}

func (user *User) CreateUser(db *mongo.Database) string {
	hashPassword(&user.Password)
	res, _ := db.Collection("users").InsertOne(context.TODO(), user)
	objectID, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return ""
	}
	id, _ := objectID.MarshalText()
	return string(id)
}

func (user *User) passwordIsValid(found *User) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(user.Password))
	return err == nil
}

func (user *User) RegisterLogin(db *mongo.Database, client *Client) (string, error) {
	var query = db.Collection("users").FindOne(context.TODO(), bson.M{"username": user.Username})

	if query.Err() != nil {
		user.CreateUser(db)
		return "", nil
	}

	var found User
	query.Decode(&found)

	if !user.passwordIsValid(&found) {
		return "", fmt.Errorf("invalid password")
	}

	id, _ := found.ID.MarshalText()
	client.user = user
	return string(id), nil
}
