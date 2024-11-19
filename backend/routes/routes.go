package routes

import (
	"database/sql"
	"backend/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Регистрация пользователя
	router.HandleFunc("/register", controllers.RegisterHandler(db)).Methods("POST")

	// Верификация пользователя
	router.HandleFunc("/verify", controllers.VerifyHandler(db)).Methods("POST")

	// Вход пользователя
	router.HandleFunc("/login", controllers.LoginHandler(db)).Methods("POST")

	return router
}
