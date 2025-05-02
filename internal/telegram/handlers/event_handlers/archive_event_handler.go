package event_handlers

import (
	"game_mill_ai_bot/internal/services/event_services/event_status"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func ArchiveEventHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /archive <id>")
	}

	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}

	response := event_status.ArchiveEventService(c.Sender().ID, localID)
	message := response_services.FormatMessage(response)

	if message != "" {
		return c.Reply(message)
	}

	return nil
}
