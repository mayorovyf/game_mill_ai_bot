package event_handlers

import (
	"game_mill_ai_bot/internal/services/event_services"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

// обработчик /delete <id>
func DeleteEventHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /delete <id>")
	}

	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}

	response := event_services.DeleteEvent(c.Sender().ID, localID)

	message := response_services.FormatMessage(response)
	if message != "" {
		return c.Reply(message)
	}

	return nil
}
