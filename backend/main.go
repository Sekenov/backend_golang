package main

import (
    "log"
    "net/http"
    "github.com/gorilla/handlers"
    "backend/config"
    "backend/routes"
)

func main() {
    // Подключение к базе данных
    db, err := config.InitDB()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }
    defer db.Close()

    // Настройка маршрутов
    router := routes.SetupRoutes(db)

    // Добавление CORS
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}),
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Запуск сервера
    log.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
