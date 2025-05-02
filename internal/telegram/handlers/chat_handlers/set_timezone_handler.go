package chat_handlers

import (
	"fmt"
	"game_mill_ai_bot/internal/services/chat_services"
	"gopkg.in/telebot.v3"
	"strings"
)

func SetTimezoneHandler(c telebot.Context) error {
	arg := strings.TrimSpace(c.Message().Payload)
	if arg == "" {
		return c.Reply("Укажите смещение в часах: /settimezone +3 или /settimezone -5")
	}

	tz, err := parseTimezone(arg)
	if err != nil {
		return c.Reply(err.Error())
	}

	if err := chat_services.SetChatTimezone(c.Chat().ID, tz); err != nil {
		return c.Reply("Ошибка установки часового пояса.")
	}

	return c.Reply(fmt.Sprintf("Часовой пояс обновлён: UTC%+d", tz))
}
