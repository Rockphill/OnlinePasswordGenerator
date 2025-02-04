package utils

import (
	"math/rand"
	"time"
)

// GeneratePassword генерирует пароль заданной длины
func GeneratePassword(length int, includeSpecial bool) string {
	// Наборы символов для генерации
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	specials := "!@#$%&*()-_=+[]{};:,.<>?/"

	// Формируем итоговый набор символов
	charset := letters
	if includeSpecial {
		charset += specials
	}

	// Инициализируем генератор случайных чисел
	rand.Seed(time.Now().UnixNano())

	// Создаем слайс для пароля
	password := make([]byte, length)

	// Генерируем пароль
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}
