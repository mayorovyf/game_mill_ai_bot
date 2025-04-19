package telegram

import (
	"fmt"
	"game_mill_ai_bot/ai"
	"game_mill_ai_bot/config"
	"gopkg.in/telebot.v3"
	"strings"
)

func AiHendler(c telebot.Context) error {
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

	c.Send("Думаю...", sendOpts)

	reply, err := ai.GetChatResponse(config.AiClient, prompt)
	if err != nil {
		return c.Send(fmt.Sprintf("Ошибка: %v", err), sendOpts)
	}

	return c.Send(reply, sendOpts)
}
