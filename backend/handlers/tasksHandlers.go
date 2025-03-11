package handlers

import (
	"backend/middleware"
	"backend/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection

func InitTaskHandlers(c *mongo.Collection) {
	taskCollection = c
}

// @SummaryGet users tasks
// @Description get all user tasks
// @Tags tasks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} map[string]interface{} "Список задач"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /tasks [get]
func GetTasks(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)

	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := taskCollection.Find(ctx, bson.M{"user_id": objID})

	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var tasks []models.Task

	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	json.NewEncoder(w).Encode(tasks)
}

// @SummaryGet Create task
// @Description Create task
// @Tags tasks
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body models.Task true "User data"
// @Success 200 {object} map[string]interface{} "Список задач"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /task [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	task := r.Context().Value(middleware.ContextTaskKey).(models.Task)

	if !ok || userID == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	task.UserID = objID

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if _, err := taskCollection.InsertOne(ctx, task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}
