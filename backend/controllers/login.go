package controllers

import (
	"database/sql"
	// Этот пакет предоставляет функции для кодирования и декодирования данных в формат JSON. Мы используем его для
	// преобразования данных из формы  json в структуру GO и наоборот
	"encoding/json"
	"net/http"
)
// LoginRequest структура которая оределяет форму данных ожидаемых в запросе для входа пользователя,
// Поля email и password будут хранить значения которые клиент отправляет на сервер для аудентивикации
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
//LoginHandler  Функция, которая возвращает обработчик  HTTP запросы. Она принимает обьект базы данных  db чтобы использовать его для работы с базой данных при проверке данных пользователя
// http.HandlerFunc - тип который реализует интерфейс http.Handler. Этот интерфес используется для обработки  http  запросов.  
func LoginHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// мы обьявляем переменную  req типа LoginRequest  которая будет хранить данные полученные в запросе
		var req LoginRequest
		// строка пытается декодировать тело апроса в  json и преорозовать его в структуру LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"message": "Invalid request"}`, http.StatusBadRequest)
			return
		}
		// storedPassword и status — переменные, которые будут хранить пароль и статус пользователя, извлеченные из базы данных.
		var storedPassword, status string
		// db.QueryRow — выполняет SQL-запрос для извлечения данных из базы. В запросе мы ищем пользователя по его email
		//.Scan(&storedPassword, &status) — эта функция сканирует результаты запроса и сохраняет их в переменные storedPassword и status. В результате мы получаем пароль пользователя и его статус (например, verified или unverified).

		err := db.QueryRow("SELECT password, status FROM users WHERE email = $1", req.Email).Scan(&storedPassword, &status)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, `{"message": "User not found"}`, http.StatusNotFound)
			} else {
				http.Error(w, `{"message": "Internal server error"}`, http.StatusInternalServerError)
			}
			return
		}

		// Если пользователь не верифицирован
		if status != "verified" {
			http.Error(w, `{"message": "User is not verified"}`, http.StatusForbidden)
			return
		}

		// Проверка пароля
		if req.Password != storedPassword {
			http.Error(w, `{"message": "Incorrect password"}`, http.StatusUnauthorized)
			return
		}

		// Успешный вход
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
	}
}
