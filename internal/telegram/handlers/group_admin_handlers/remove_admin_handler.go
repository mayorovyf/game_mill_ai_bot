package group_admin_handlers

import (
	"game_mill_ai_bot/internal/services/group_services/admin_services"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func RemoveAdminHandler(c telebot.Context) error {
	arg := strings.TrimSpace(c.Message().Payload)
	userID, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return c.Reply("user_id должен быть числом")
	}

	resp := admin_services.RemoveAdmin(c.Chat().ID, userID)
	msg := response_services.FormatMessage(resp)
	if msg != "" {
		return c.Reply(msg)
	}
	return nil
}
