package routes

import (
	"backend/controllers"
	"database/sql"
	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Регистрация пользователя
	router.HandleFunc("/register", controllers.RegisterHandler(db)).Methods("POST")

	// Верификация пользователя
	router.HandleFunc("/verify", controllers.VerifyHandler(db)).Methods("POST")

	return router
}
