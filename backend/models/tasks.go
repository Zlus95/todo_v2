package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title  string             `json:"title" bson:"title"`
	Status string             `json:"status" bson:"status"`
}

type CreateTask struct {
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title  string             `json:"title" bson:"title"`
	Status string             `json:"status" bson:"status"`
}
