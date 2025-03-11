package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title  string             `json:"title" bson:"title"`
	Status string             `json:"status" bson:"status"`
}

type CreateTask struct {
	Title  string `json:"title" bson:"title"`
	Status string `json:"status" bson:"status"`
}

type UpdateTask struct {
	Title  string `json:"title" bson:"title"`
	Status string `json:"status" bson:"status"`
}

type TaskResponse struct {
	ID     primitive.ObjectID `json:"id,omitempty"`
	Title  string             `json:"title"`
	Status string             `json:"status"`
}
