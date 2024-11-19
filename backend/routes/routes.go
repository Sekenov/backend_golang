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

	return router
}
