package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username"`
	Content  string             `bson:"content"`
}
