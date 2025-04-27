package main_handlers

import (
	"game_mill_ai_bot/internal/services/user_services"
	"gopkg.in/telebot.v3"
)

func StartHandler(c telebot.Context) error {
	if c.Chat().Type != telebot.ChatPrivate {
		return c.Reply("Эта команда работает только в @" + c.Bot().Me.Username)
	}

	response := user_services.CreateUser(c.Sender())

	if response.VisibleToUser {
		return c.Reply(response.Description)
	}

	return nil
}
