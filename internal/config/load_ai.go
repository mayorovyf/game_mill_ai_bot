// internal/config/load_ai.go
package config

import (
	"github.com/sashabaranov/go-openai"
	"log"
	"os"
)

var AiClient *openai.Client

// Инициализация OpenAI клиента
func LoadAIConfig() {
	openaiKey := os.Getenv("OPENAI_API_KEY")
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY не найден в .env")
	}

	AiClient = openai.NewClient(openaiKey)
	log.Println("OpenAI клиент успешно инициализирован")
}
