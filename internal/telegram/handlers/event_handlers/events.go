package event_handlers

import (
	"fmt"
	"game_mill_ai_bot/internal/services/event_services"
	"gopkg.in/telebot.v3"
)

func ListEventsHandler(c telebot.Context) error {
	events, err := event_services.ListUserEvents(c.Sender().ID)
	if err != nil {
		return c.Reply("Ошибка получения списка событий.")
	}
	if len(events) == 0 {
		return c.Reply("У вас нет событий.")
	}
	out := ""
	for _, ev := range events {
		out += fmt.Sprintf("ID: %d [%s] %s\n", ev.LocalID, ev.Status, ev.Title)
	}
	return c.Reply(out)
}
