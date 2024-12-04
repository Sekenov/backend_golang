package routes

import (
	// пает для работы с базза данных sql
	"database/sql"
	// для апишок
	"backend/controllers"
	
	"github.com/gorilla/mux"
)
// функция, которая настраивает все маршруты для нашего API.
//db *sql.DB — мы передаем в эту функцию объект соединения с базой данных, который будет использоваться обработчиками запросов для работы с данными в базе.
func SetupRoutes(db *sql.DB) *mux.Router {
	//  создаем новый маршрутизатор с помощью библиотеки mux. Он будет использоваться для обработки всех входящих HTTP-запросов.
	router := mux.NewRouter()

	// Регистрация пользователя
	router.HandleFunc("/register", controllers.RegisterHandler(db)).Methods("POST")

	// Верификация пользователя
	router.HandleFunc("/verify", controllers.VerifyHandler(db)).Methods("POST")

	// Вход пользователя
	router.HandleFunc("/login", controllers.LoginHandler(db)).Methods("POST")

	return router
}
