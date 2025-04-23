// internal/config/load_env.go
package config

import (
	"github.com/joho/godotenv"
	"log"
)

// загрузка .env файла
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}
}
