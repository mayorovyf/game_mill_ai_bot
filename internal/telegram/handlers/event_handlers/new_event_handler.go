package event_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"game_mill_ai_bot/internal/services/event_services/draft"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
)

func NewEventHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	response = draft.CreateDraft(c.Sender().ID, c.Chat().ID, nil)

	message := response_services.FormatMessage(response)
	if message != "" {
		return c.Reply(message)
	}

	return nil
}
