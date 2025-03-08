package main

import (
	"backend/handlers"
	"context"
	"log"
	"net/http"

	_ "backend/docs"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title User Registration API
// @version 1.0
// @description API для регистрации пользователей
// @host localhost:8080
// @BasePath /
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

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

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
