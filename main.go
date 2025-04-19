package main

import (
	"game_mill_ai_bot/config"
	"game_mill_ai_bot/db"
	"game_mill_ai_bot/telegram"
	"log"
	"os"
	"time"

	"gopkg.in/telebot.v3"
)

func main() {

	config.LoadEnv()
	config.LoadAIConfig()
	db.ConnectDB()

	botApiKey := os.Getenv("TG_BOT_API_KEY")

	if botApiKey == "" {
		log.Fatal("TG_BOT_API_KEY не найден в .env")
	}

	// Настройки бота
	pref := telebot.Settings{
		Token:  botApiKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	bot.Handle("/start", telegram.StartHandler)
	bot.Handle("/ai", telegram.AiHendler)
	bot.Handle("/profile", telegram.ProfileHandler)
	bot.Handle("/add", telegram.AddCloudletsHandler)

	log.Println("Бот запущен...")
	bot.Start()
}
