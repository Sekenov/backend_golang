package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type VerifyRequest struct {
	Email             string `json:"email"`
	VerificationCode  string `json:"verification_code"`
}

func VerifyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req VerifyRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
			return
		}

		// Проверка кода и статуса пользователя
		var storedCode, status string
		err := db.QueryRow("SELECT verification_code, status FROM users WHERE email = $1", req.Email).Scan(&storedCode, &status)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, `{"message": "Internal server error"}`, http.StatusInternalServerError)
			}
			return
		}

		// Если код неверный
		if storedCode != req.VerificationCode {
			http.Error(w, `{"message": "Invalid verification code"}`, http.StatusBadRequest)
			return
		}

		// Если код верный, обновляем статус на "verified"
		_, err = db.Exec("UPDATE users SET status = 'verified' WHERE email = $1", req.Email)
		if err != nil {
			http.Error(w, `{"message": "Failed to update user status"}`, http.StatusInternalServerError)
			return
		}

		// Ответ при успешной верификации
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User verified successfully"})
	}
}
