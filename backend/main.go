package main

import (
	"backend/handlers"
	"backend/middleware"
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Подключение к MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	tasksCollection := client.Database("todolist").Collection("tasks")
	userCollection := client.Database("todolist").Collection("users")

	// Инициализация маршрутизатора
	r := mux.NewRouter()
	handlers.InitTaskHandlers(tasksCollection)
	handlers.InitAuthHandlers(userCollection)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	r.HandleFunc("/register", middleware.ValidRegister(handlers.Register)).Methods("POST")
	r.HandleFunc("/login", middleware.ValidLogin(handlers.Login)).Methods("POST")
	r.Handle("/tasks", middleware.AuthMiddleware(http.HandlerFunc(handlers.GetTasks))).Methods("GET")
	r.Handle("/task", middleware.AuthMiddleware(
		middleware.ValidTask(
			http.HandlerFunc(handlers.CreateTask),
		),
	)).Methods("POST")
	r.Handle("/task/{id}", middleware.AuthMiddleware(
		middleware.ValidTask(
			http.HandlerFunc(handlers.UpdateTask),
		),
	)).Methods("PATCH")
	r.Handle("/task/{id}", middleware.AuthMiddleware(http.HandlerFunc(handlers.DeleteTask))).Methods("DELETE")

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
