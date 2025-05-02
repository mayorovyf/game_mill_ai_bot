package event_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"game_mill_ai_bot/internal/services/event_services/edit"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SetHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	args := strings.Fields(c.Message().Payload)
	if len(args) < 3 {
		return c.Reply("Формат: /set <поле> <id> <значение>")
	}

	field := args[0]
	localID, err := strconv.Atoi(args[1])
	if err != nil {
		return c.Reply("ID события должен быть числом")
	}

	value := strings.Join(args[2:], " ")

	response = edit.UpdateEventField(c.Sender().ID, localID, field, value)
	message := response_services.FormatMessage(response)

	if message != "" {
		return c.Reply(message)
	}

	return nil
}
