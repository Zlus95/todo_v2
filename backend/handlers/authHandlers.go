package handlers

import (
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitAuthHandlers(c *mongo.Collection) {
	userCollection = c
}

// @Summary New user registration
// @Description Registers a new user in the system
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]string "The user has been successfully registered"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 409 {object} map[string]string "User already exists"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	foundUser := bson.M{"email": user.Email}
	if err := userCollection.FindOne(ctx, foundUser).Err(); err == nil {
		log.Printf("User with email %s already exists", user.Email)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Error checking user existence: %v", err)
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		return
	}

	if user.Email == "" || !strings.Contains(user.Email, "@") {
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	if _, err := userCollection.InsertOne(ctx, user); err != nil {
		log.Printf("Error inserting user: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// @Summary User authorization
// @Description Authorization
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 201 {object} map[string]string "The user has been successfully authorized."
// @Failure 401 {object} map[string]string "Invalid Password"
// @Failure 401 {object} map[string]string "Invalid Email"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var currentUser models.User
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&currentUser); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("User not found: %v", err)
			http.Error(w, "Invalid Email", http.StatusUnauthorized)
			return
		}
		log.Printf("Error finding user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(user.Password)); err != nil {
		log.Printf("Invalid password for user %s: %v", user.Email, err)
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(currentUser.ID.Hex())
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
