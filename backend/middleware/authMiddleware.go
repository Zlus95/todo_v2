package middleware

import (
	"backend/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func ValidLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserLogin

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if user.Email == "" || !strings.Contains(user.Email, "@") {
			http.Error(w, "Invalid email", http.StatusBadRequest)
			return
		}

		if user.Password == "" {
			http.Error(w, "Password are required", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func ValidRegister(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserRegister

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Printf("Error decoding request body: %v", err)
			http.Error(w, "Invalid request body", http.StatusBadRequest)
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

		if len(user.Name) < 2 {
			http.Error(w, "Name must be at least 2 characters", http.StatusBadRequest)
			return
		}
		if len(user.LastName) < 2 {
			http.Error(w, "Last Name must be at least 2 characters", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
