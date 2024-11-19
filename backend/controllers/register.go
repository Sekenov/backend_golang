package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"backend/utils"
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
			http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
			return
		}

		// Генерация 6-значного кода
		code := utils.GenerateVerificationCode()

		// Отправка письма с кодом
		err := utils.SendVerificationCode(req.Email, code)
		if err != nil {
			http.Error(w, `{"message": "Failed to send verification code"}`, http.StatusInternalServerError)
			fmt.Println("Error sending email:", err)
			return
		}

		// Сохранение кода в базу данных для последующей проверки
		_, err = db.Exec("INSERT INTO users (name, last_name, email, password, verification_code) VALUES ($1, $2, $3, $4, $5)",
			req.Name, req.LastName, req.Email, req.Password, code)
		if err != nil {
			http.Error(w, `{"message": "Failed to save user"}`, http.StatusInternalServerError)
			fmt.Println("Error saving user to database:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully. Check your email for the verification code."})
	}
}
