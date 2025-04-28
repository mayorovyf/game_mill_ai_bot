package chat_services

import "gopkg.in/telebot.v3"

// проверям является ли чат группой
func IsGroupChat(c telebot.Context) bool {
	return c.Chat().Type == telebot.ChatGroup
}
