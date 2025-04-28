package chat_services

import "gopkg.in/telebot.v3"

// проверяем является ли этот чат личным
func IsPrivateChat(c telebot.Context) bool {
	return c.Chat().Type == telebot.ChatPrivate
}
