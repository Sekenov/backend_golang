//в GO main.go является основным файлом, всегда приложения начинается с пакета main, Это указывает на то, что файл содеержит точку входа в приложение
package main

// Импортируем необходимые нам пакеты
import (
    // для логирования ошибок и просто других сообщение
    "log"
    //этот пакет нам помогает работать с http запросами
    "net/http"
    // это пакет из популярной библиотеки gorilla. В нем есть утилиты для работы с HTTP-обработчиками, например, для обработки CORS (Cross-Origin Resource Sharing), чтобы разрешить запросы с других доменов.
    "github.com/gorilla/handlers"
    // пакет, который используетсся для инициализации базы данных
    "backend/config"
    // пакет, который отвечает за наши маршруты
    "backend/routes"
)
// когда мы запускаем проект именно эта функция отвечает за всю работу
func main() {
    // ызываем функцию InitDB из пакета config, которая отвечает за установление соединения с базой данных. Если соединение не удается, возвращается ошибка.
    db, err := config.InitDB()
    // если ошибка при подключения к базе данных то выводим текст об этом,
    if err != nil {
        // лог фатал завершить процесс после вывода ошибки
        log.Fatal("Failed to connect to the database:", err)
    }
    //откладывает выполнение функции до завершения функции main. В данном случае она генерирует что соединение с базой данных будет закрыто в конце работы программы.
    defer db.Close()

    // создааем маршрут для нашего api, SetupRoutes из пакета routes возвращает маршруты, настроенные для работы с базой данных
    router := routes.SetupRoutes(db)
    // njddd
    // Добавление CORS для того чтобы наш фронт который на 3000 мог работать с бэкенд который на 8080
    // handlers.CORS это функция из библиотеки gorilla/handlers, которая позволяет настроить CORS.
    corsHandler := handlers.CORS(
        // говорим то что фронт который в 3000 сервере может работать с бэкендом
        handlers.AllowedOrigins([]string{"http://localhost:3000"}),
        // разрешаем использовать методы HTTP, такие
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
        // разрешаем отправлять заголовки
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
    )

    // Запуск сервера
    //выводим в лог сообщение, что сервер успешно запущен и слушает порт 8080.
    log.Println("Server is running on port 8080...")
    // дает ошибку какую то при ошибке
    log.Fatal(http.ListenAndServe(":8080", corsHandler(router)))
}
