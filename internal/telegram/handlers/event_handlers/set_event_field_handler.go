package event_handlers

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"gopkg.in/telebot.v3"
	"strings"
	"time"
)

func SetEventFieldHandler(c telebot.Context) error {
	args := c.Args()
	if len(args) < 2 {
		return c.Reply("Используй:\n/set_event <ID> tag новое_значение\n" +
			"Пример: /set_event 1234567890 title МояИгра")
	}
	eventID := args[0]
	rest := strings.Join(args[1:], " ")
	parts := strings.SplitN(rest, " ", 2)
	if len(parts) < 2 {
		return c.Reply("Нужно указать тег и значение, например:\n" +
			"/set_event " + eventID + " description1 Краткое описание")
	}
	tag, value := parts[0], parts[1]

	evt, err := r_event.GetEventById(eventID)
	if err != nil || evt == nil {
		return c.Reply("Событие не найдено")
	}

	switch tag {
	case "title":
		evt.Title = value
	case "description1":
		evt.Description1 = value
	case "description2":
		evt.Description2 = value
	case "type_id":
		evt.TypeID = value
	case "type":
		evt.Type = value
	case "date":
		// ожидаем формат YYYY-MM-DD_HH:MM
		dateStr := strings.ReplaceAll(value, "_", " ")
		t, err := time.Parse("2006-01-02 15:04", dateStr)
		if err != nil {
			return c.Reply("Неверный формат даты. Используй: YYYY-MM-DD_HH:MM")
		}
		evt.Date = t
	default:
		return c.Reply("Неизвестный тег: " + tag)
	}

	if err := r_event.UpdateEvent(evt); err != nil {
		return c.Reply("Ошибка при обновлении: " + err.Error())
	}
	return c.Reply("Поле `" + tag + "` обновлено на: " + value)
}
