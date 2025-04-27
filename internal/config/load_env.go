// internal/config/load_env.go
package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// загрузка .env файла
func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Файл .env не найден")
	}

	mode := Mode(os.Getenv("MODE"))

	switch mode {
	case DevMode, TestMode, ProdMode:
		CurrentMode = mode
	default:
		CurrentMode = ProdMode
	}

	log.Println("Текущий режим -", CurrentMode)
}
