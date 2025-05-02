// internal/telegram/handlers/main_handlers/start_handler.go
package main_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"game_mill_ai_bot/internal/services/response_services"
	"game_mill_ai_bot/internal/services/user_services"
	"gopkg.in/telebot.v3"
)

// обработчик start
func StartHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	if c.Chat().Type != telebot.ChatPrivate {
		return c.Reply("Эта команда работает только в @" + c.Bot().Me.Username)
	}

	response = user_services.CreateUser(c.Sender())

	message := response_services.FormatMessage(response)
	if message != "" {
		return c.Reply(message)
	}

	return nil
}
