package event_handlers

import (
	"fmt"
	"game_mill_ai_bot/internal/services/event_services"
	"gopkg.in/telebot.v3"
)

func NewEventHandler(c telebot.Context) error {
	event, err := event_services.CreateDraft(
		c.Sender().ID,
		c.Chat().ID,
		nil, // можно добавить поддержку topicID если нужно
	)
	if err != nil {
		return c.Reply("Ошибка создания события: " + err.Error())
	}
	return c.Reply(fmt.Sprintf(
		"Создан черновик события с ID: %d\nДля заполнения используйте /set <поле> %d <значение>",
		event.LocalID, event.LocalID,
	))
}
