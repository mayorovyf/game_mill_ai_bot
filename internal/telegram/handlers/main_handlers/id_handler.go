package main_handlers

import (
	"fmt"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func IDHandler(c telebot.Context) error {
	// 1. Reply: если команда дана в ответ на сообщение
	if c.Message().IsReply() && c.Message().ReplyTo.Sender != nil {
		target := c.Message().ReplyTo.Sender
		return c.Reply(fmt.Sprintf("ID пользователя: %d", target.ID))
	}

	// 2. Аргумент: username или id
	arg := strings.TrimSpace(c.Message().Payload)
	if arg != "" {
		// Попробуем как ID
		if id, err := strconv.ParseInt(arg, 10, 64); err == nil {
			return c.Reply(fmt.Sprintf("ID пользователя: %d", id))
		}

		// Попробуем как username
		if !strings.HasPrefix(arg, "@") {
			arg = "@" + arg
		}

		user, err := c.Bot().ChatByUsername(arg)
		if err != nil || user == nil {
			return c.Reply("Не удалось найти пользователя.")
		}

		return c.Reply(fmt.Sprintf("ID пользователя @%s: %d", user.Username, user.ID))
	}

	// 3. По умолчанию — показать свой ID
	return c.Reply(fmt.Sprintf("Ваш ID: %d", c.Sender().ID))
}
