package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Проверяем обязательные поля
		if req.Name == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Сохраняем пользователя в базе данных
		_, err := db.Exec(
			"INSERT INTO users (name, last_name, email, password) VALUES ($1, $2, $3, $4)",
			req.Name, req.LastName, req.Email, req.Password,
		)
		if err != nil {
			http.Error(w, "Failed to save user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
	}
}
