package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
			return
		}

		var storedPassword, status string
		err := db.QueryRow("SELECT password, status FROM users WHERE email = $1", req.Email).Scan(&storedPassword, &status)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, `{"message": "Internal server error"}`, http.StatusInternalServerError)
			}
			return
		}

		// Если пользователь не верифицирован
		if status != "verified" {
			http.Error(w, `{"message": "User is not verified"}`, http.StatusForbidden)
			return
		}

		// Проверка пароля
		if req.Password != storedPassword {
			http.Error(w, `{"message": "Incorrect password"}`, http.StatusUnauthorized)
			return
		}

		// Успешный вход
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	}
}
