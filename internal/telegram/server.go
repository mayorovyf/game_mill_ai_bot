package telegram

import (
	"gopkg.in/telebot.v3"
	"log"
	"os"
	"time"
)

var bot *telebot.Bot

func LoadTgBot() {
	botApiKey := os.Getenv("TG_BOT_API_KEY")

	if botApiKey == "" {
		log.Fatal("TG_BOT_API_KEY не найден в .env")
	}

	// Настройки бота
	pref := telebot.Settings{
		Token:  botApiKey,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	var err error
	bot, err = telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	LoadTgRoutes()

	log.Println("Бот запущен...")
	bot.Start()
}
