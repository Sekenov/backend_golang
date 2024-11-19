package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type VerifyRequest struct {
	Email            string `json:"email"`
	VerificationCode string `json:"verification_code"`
}

func VerifyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req VerifyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
			return
		}

		// Проверяем, что код в базе данных совпадает с тем, что прислал пользователь
		var storedCode string
		err := db.QueryRow("SELECT verification_code FROM users WHERE email = $1", req.Email).Scan(&storedCode)
		if err != nil {
			http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
			fmt.Println("Error fetching user:", err)
			return
		}

		// Сравниваем коды
		if req.VerificationCode != storedCode {
			http.Error(w, `{"message": "Invalid verification code"}`, http.StatusBadRequest)
			return
		}

		// Обновляем статус пользователя в базе данных на "verified"
		_, err = db.Exec("UPDATE users SET status = 'verified' WHERE email = $1", req.Email)
		if err != nil {
			http.Error(w, `{"message": "Failed to update user status"}`, http.StatusInternalServerError)
			fmt.Println("Error updating user status:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Verification successful!"})
	}
}
