package handlers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"web/utils"
)

// Middleware для логирования
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Логируем информацию о запросе
		log.Printf(
			"Started %s %s from %s",
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
		)

		next.ServeHTTP(w, r)

		// Логируем время выполнения
		log.Printf(
			"Completed %s %s in %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	}
}

// HomeHandler обрабатывает главную страницу
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("404 for path: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")

	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// PasswordHandler обрабатывает запросы на генерацию пароля
func PasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lengthStr := r.URL.Query().Get("length")
	includeSpecialStr := r.URL.Query().Get("special")

	length := 8
	includeSpecial := false

	if lengthStr != "" {
		if l, err := strconv.Atoi(lengthStr); err == nil && l >= 8 && l <= 12 {
			length = l
		} else {
			http.Error(w, "Invalid length parameter", http.StatusBadRequest)
			return
		}
	}

	if includeSpecialStr == "true" {
		includeSpecial = true
	}

	password := utils.GeneratePassword(length, includeSpecial)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Write([]byte(password))
}
