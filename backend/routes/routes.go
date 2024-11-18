package routes

import (
	"github.com/gorilla/mux"
	"backend/controllers"
	"backend/config"
)

func SetupRoutes(router *mux.Router) {
	// Роут для регистрации
	router.HandleFunc("/register", controllers.RegisterHandler(config.DB)).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.LoginHandler(config.DB)).Methods("POST", "OPTIONS")

}
