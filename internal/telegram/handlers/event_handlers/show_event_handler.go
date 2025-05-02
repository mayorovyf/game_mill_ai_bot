package event_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"game_mill_ai_bot/internal/services/event_services"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func ShowEventHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /showevent <id>")
	}

	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}

	response = event_services.GetEventCard1(c.Sender().ID, localID)
	message := response_services.FormatMessage(response)

	if message != "" {
		return c.Reply(message)
	}

	return nil
}
