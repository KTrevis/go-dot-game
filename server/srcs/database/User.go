package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

func (user *User) UsernameTaken(db *mongo.Database) bool {
	var found = db.Collection("users").FindOne(context.TODO(), bson.M{"username": user.Username})
	return found.Err() == nil
}

func (user *User) CreateUser(db *mongo.Database) {
	db.Collection("users").InsertOne(context.TODO(), user)
}
