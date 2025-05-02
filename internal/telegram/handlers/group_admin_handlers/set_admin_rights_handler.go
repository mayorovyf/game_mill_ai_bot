package group_admin_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/group_services/admin_services"
	"game_mill_ai_bot/internal/services/response_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SetAdminRightsHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 2 {
		return c.Reply("Пример: /setadminrights <user_id> can_edit=true can_archive=false")
	}

	userID, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		return c.Reply("user_id должен быть числом")
	}

	priv := models.AdminPrivileges{}
	for _, pair := range args[1:] {
		if strings.Contains(pair, "=") {
			parts := strings.SplitN(pair, "=", 2)
			key := strings.ToLower(parts[0])
			val := strings.ToLower(parts[1]) == "true"

			switch key {
			case "can_edit":
				priv.CanEditEvents = val
			case "can_archive":
				priv.CanArchive = val
			case "can_manage":
				priv.CanManageMembers = val
			case "can_delete":
				priv.CanDeleteMessages = val
			default:
				// игнор неизвестных ключей
			}
		}
	}

	response := admin_services.SetAdminRights(c.Chat().ID, userID, priv)
	message := response_services.FormatMessage(response)
	if message != "" {
		return c.Reply(message)
	}
	return nil
}
