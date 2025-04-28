package event_handlers

import (
	"fmt"
	"game_mill_ai_bot/internal/services/event_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func ShowEventHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /showevent <id>")
	}
	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}
	ev, err := event_services.GetEventCard(c.Sender().ID, localID)
	if err != nil {
		return c.Reply("Событие не найдено.")
	}
	reminders := strings.Trim(strings.Replace(fmt.Sprint(ev.ReminderMins), " ", ", ", -1), "[]")
	out := fmt.Sprintf(
		"Событие: %s\nВремя: %s UTC\nОписание: %s\nНапоминания: %s\nПодписчики: %d",
		ev.Title, ev.StartTime.Format("2006-01-02 15:04"), ev.Description, reminders, len(ev.Subscribers),
	)
	return c.Reply(out)
}
