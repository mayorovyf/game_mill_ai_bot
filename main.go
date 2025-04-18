package main

import (
	"fmt"
	"game_mill_ai_bot/ai"
	"game_mill_ai_bot/config"
	"game_mill_ai_bot/db"
	"game_mill_ai_bot/telegram"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/telebot.v3"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	config.LoadEnv()
	db.ConnectDB()

	botApiKey := os.Getenv("TG_BOT_API_KEY")
	openaiKey := os.Getenv("OPENAI_API_KEY")

	if botApiKey == "" {
		log.Fatal("TG_BOT_API_KEY не найден в .env")
	}
	if openaiKey == "" {
		log.Fatal("OPENAI_API_KEY не найден в .env")
	}

	// Инициализируем OpenAI клиент
	aiClient := openai.NewClient(openaiKey)

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

	bot.Handle("/ai", func(c telebot.Context) error {
		message := c.Message()
		prompt := strings.TrimSpace(message.Payload)

		if c.Chat().Type != telebot.ChatSuperGroup {
			return c.Send("Бот работает только в супергруппах")
		}
		if message.ThreadID == 0 {
			return c.Send("Пожалуйста, используй команду в топике")
		}
		if prompt == "" {
			return c.Send("Пожалуйста, укажи запрос после команды, например:\n`/ai Что такое черная дыра?`", &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		}

		sendOpts := &telebot.SendOptions{
			ThreadID:  message.ThreadID,
			ParseMode: telebot.ModeMarkdown,
		}

		c.Send("Думаю... 🤖", sendOpts)

		reply, err := ai.GetChatResponse(aiClient, prompt)
		if err != nil {
			return c.Send(fmt.Sprintf("Ошибка: %v", err), sendOpts)
		}

		return c.Send(reply, sendOpts)
	})

	log.Println("Бот запущен...")
	bot.Start()
}
