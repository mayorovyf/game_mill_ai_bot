package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"

	openai "github.com/sashabaranov/go-openai"
)

// Функция запроса к OpenAI
func getChatResponse(client *openai.Client, input string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    "user",
					Content: input,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func main() {

	// Загрузка конфигурации
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

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

	bot.Handle("/start", func(c telebot.Context) error {
		chat := c.Chat()
		message := c.Message()
		topicSendOptions := &telebot.SendOptions{
			ThreadID: message.ThreadID,
		}

		// Сохраняем цепочку ошибок
		var err error

		if chat.Type == telebot.ChatSuperGroup {
			if e := c.Send("Сообщение из супергруппы"); e != nil {
				err = e
			}
		}

		if message.ThreadID != 0 {
			if e := c.Send("Сообщение из топика"); e != nil {
				err = e
			}

			if e := c.Send("ID топика: " + strconv.Itoa(message.ThreadID)); e != nil {
				err = e
			}

			bot.Send(chat, "Сообщение отправлено в топик откуда получена команда", topicSendOptions)
		}

		return err
	})

	bot.Handle("/ai", func(c telebot.Context) error {
		prompt := strings.TrimSpace(c.Message().Payload)

		if prompt == "" {
			return c.Send("Пожалуйста, укажи запрос после команды, например:\n`/ai Что такое черная дыра?`", &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
		}

		c.Send("Думаю... 🤖")

		reply, err := getChatResponse(aiClient, prompt)
		if err != nil {
			return c.Send(fmt.Sprintf("Ошибка: %v", err))
		}

		return c.Send(reply)
	})

	log.Println("Бот запущен...")
	bot.Start()
}
