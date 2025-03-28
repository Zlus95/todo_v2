package handlers

import (
	"backend/models"
	"backend/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
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
// @Param user body models.UserRegister true "User data"
// @Success 201 {object} map[string]string "The user has been successfully registered"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 409 {object} map[string]string "User already exists"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /register [post]
func Register(w http.ResponseWriter, r *http.Request) {
	userReg, ok := r.Context().Value("user").(models.UserRegister)

	if !ok {
		log.Printf("Error getting user from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	foundUser := bson.M{"email": userReg.Email}
	if err := userCollection.FindOne(ctx, foundUser).Err(); err == nil {
		log.Printf("User with email %s already exists", userReg.Email)
		http.Error(w, "User already exists", http.StatusConflict)
		return
	} else if err != mongo.ErrNoDocuments {
		log.Printf("Error checking user existence: %v", err)
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReg.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Failed to process password", http.StatusInternalServerError)
		return
	}

	userReg.Password = string(hashedPassword)

	if _, err := userCollection.InsertOne(ctx, userReg); err != nil {
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
// @Param user body models.UserLogin true "User data"
// @Success 201 {object} map[string]string "The user has been successfully authorized."
// @Failure 401 {object} map[string]string "Invalid Password"
// @Failure 401 {object} map[string]string "Invalid Email"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	userLogin, ok := r.Context().Value("user").(models.UserLogin)

	if !ok {
		log.Printf("Error getting user from context")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var currentUser models.User
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	if err := userCollection.FindOne(ctx, bson.M{"email": userLogin.Email}).Decode(&currentUser); err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("User not found: %v", err)
			http.Error(w, "Invalid Email", http.StatusUnauthorized)
			return
		}
		log.Printf("Error finding user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

	}

	if err := bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(userLogin.Password)); err != nil {
		log.Printf("Invalid password for user %s: %v", userLogin.Email, err)
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateJWT(currentUser.ID.Hex())
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
