package middleware

import (
	"backend/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func ValidLigin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

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
