package main

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/telegram"
)

func main() {
	config.LoadEnv()
	config.LoadAIConfig()
	db.ConnectDB()
	telegram.LoadTgBot()
}
