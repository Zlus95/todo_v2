package middleware

import (
	"backend/models"
	"context"
	"encoding/json"
	"net/http"
)

func ValidTask(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if task.Title == "" || len(task.Title) > 20 {
			http.Error(w, "Invalid title", http.StatusBadRequest)
			return
		}

		validStatuses := map[string]bool{"to do": true, "doing": true, "done": true}
		if task.Status != "" && !validStatuses[task.Status] {
			http.Error(w, "Invalid status. Allowed values: to do, doing, done", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), ContextTaskKey, task)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
