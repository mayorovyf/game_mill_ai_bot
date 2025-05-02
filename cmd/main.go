package main

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/telegram"
	"log"
	"os"
)

func main() {
	config.LoadEnv()
	config.LoadAIConfig()
	db.ConnectDB()
	log.SetOutput(os.Stdout)
	telegram.LoadTgBot()
}
