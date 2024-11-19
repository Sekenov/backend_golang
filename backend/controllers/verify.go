package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

func VerifyHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var req struct {
			Email string `json:"email"`
			Code  string `json:"code"`
		}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Failed to parse request body", http.StatusBadRequest)
			return
		}

		var storedCode string
		err = db.QueryRow("SELECT verification_code FROM users WHERE email = $1", req.Email).Scan(&storedCode)
		if err != nil {
			http.Error(w, "Invalid email or code", http.StatusUnauthorized)
			return
		}

		if req.Code != storedCode {
			http.Error(w, "Invalid code", http.StatusUnauthorized)
			return
		}

		_, err = db.Exec("UPDATE users SET verification_code = NULL WHERE email = $1", req.Email)
		if err != nil {
			http.Error(w, "Failed to verify user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User verified successfully"})
	}
}
