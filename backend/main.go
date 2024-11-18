package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"backend/config"
	"backend/routes"
)

// Middleware для разрешения CORS
func enableCORS(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}


func main() {
	// Инициализация базы данных
	config.InitDB()

	// Создаём роутер
	router := mux.NewRouter()

	// Применяем middleware для CORS
	router.Use(enableCORS)

	// Настраиваем маршруты
	routes.SetupRoutes(router)

	// Запускаем сервер
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
