package models



type User struct {
	Name             string `json:"name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	VerificationCode string `json:"verification_code"`
}
