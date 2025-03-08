package main

import (
	"backend/handlers"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Подключение к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// tasksCollection := client.Database("todolist").Collection("tasks")
	userCollection := client.Database("todolist").Collection("users")

	// Инициализация маршрутизатора
	r := mux.NewRouter()
	// handlers.InitTaskHandlers(tasksCollection)
	handlers.InitAuthHandlers(userCollection)

	r.HandleFunc("/register", handlers.Register).Methods("Post")

	// Настройка CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Оберните маршрутизатор в CORS
	handler := c.Handler(r)

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
