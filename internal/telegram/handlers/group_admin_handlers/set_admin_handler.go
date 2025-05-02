package group_admin_handlers

import (
	"game_mill_ai_bot/internal/services/group_services/admin_services"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SetAdminHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 2 {
		return c.Reply("Пример: /setadmin <user_id> <звание>")
	}

	userID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return c.Reply("user_id должен быть числом")
	}

	title := strings.Join(args[1:], " ")
	response := admin_services.AssignAdmin(c.Chat().ID, userID, title)

	message := response_services.FormatMessage(response)
	if message != "" {
		return c.Reply(message)
	}
	return nil
}
