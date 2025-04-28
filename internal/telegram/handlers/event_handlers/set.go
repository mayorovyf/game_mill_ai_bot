package event_handlers

import (
	"game_mill_ai_bot/internal/services/event_services"
	"game_mill_ai_bot/internal/utils"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SetHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 3 {
		return c.Reply("Формат: /set <поле> <id> <значение>")
	}
	field := args[0]
	localID, err := strconv.Atoi(args[1])
	if err != nil {
		return c.Reply("ID события должен быть числом")
	}
	value := strings.Join(args[2:], " ")

	switch field {
	case "title":
		return setTitleHandler(c, localID, value)
	case "description":
		return setDescriptionHandler(c, localID, value)
	case "time":
		return setTimeHandler(c, localID, value)
	case "reminder":
		return setReminderHandler(c, localID, value)
	case "topic":
		return setTopicHandler(c, localID, value)
	default:
		return c.Reply("Неизвестное поле: " + field)
	}
}

func setTitleHandler(c telebot.Context, localID int, value string) error {
	if err := event_services.UpdateTitle(c.Sender().ID, localID, value); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Название обновлено.")
}

func setDescriptionHandler(c telebot.Context, localID int, value string) error {
	if err := event_services.UpdateDescription(c.Sender().ID, localID, value); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Описание обновлено.")
}

func setTimeHandler(c telebot.Context, localID int, value string) error {
	t, err := utils.ParseEventTime(value)
	if err != nil {
		return c.Reply(err.Error())
	}
	if err := event_services.UpdateTime(c.Sender().ID, localID, t); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Время обновлено (UTC): " + t.Format("2006-01-02 15:04"))
}

func setReminderHandler(c telebot.Context, localID int, value string) error {
	mins, err := strconv.Atoi(value)
	if err != nil {
		return c.Reply("Минуты до события должны быть числом")
	}
	if err := event_services.AddReminder(c.Sender().ID, localID, mins); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Напоминание добавлено.")
}

func setTopicHandler(c telebot.Context, localID int, value string) error {
	topicID, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return c.Reply("TopicID должен быть числом")
	}
	if err := event_services.SetTopic(c.Sender().ID, localID, topicID); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Топик обновлён.")
}
