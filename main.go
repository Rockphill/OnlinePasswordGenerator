package main

import (
	"/web/config"
	"/web/handlers"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Получаем порт из переменной окружения или используем порт по умолчанию
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // порт по умолчанию
	}

	// Инициализация конфигурации
	cfg := config.NewConfig()

	// Логирование запуска сервера
	log.Printf("Starting server on port %s...", port)
	log.Printf("Server configuration: %+v", cfg)

	// Настройка обработчиков с middleware для логирования
	http.Handle("/", handlers.LoggingMiddleware(handlers.HomeHandler))
	http.Handle("/generate", handlers.LoggingMiddleware(handlers.PasswordHandler))

	// Запуск сервера
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
