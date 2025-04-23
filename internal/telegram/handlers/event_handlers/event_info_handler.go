package event_handlers

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func EventInfoHandler(c telebot.Context) error {
	args := c.Args()
	if len(args) < 1 {
		return c.Reply("Используй:\n/get_event <ID>\nПример: /get_event 1234567890")
	}
	eventID := args[0]

	evt, err := r_event.GetEventById(eventID)
	if err != nil || evt == nil {
		return c.Reply("Событие не найдено.")
	}

	var sb strings.Builder
	sb.WriteString("ℹ️ Информация о событии (ID: <code>" + eventID + "</code>)\n\n")
	sb.WriteString("Title: " + evt.Title + "\n")
	sb.WriteString("Description1: " + evt.Description1 + "\n")
	sb.WriteString("Description2: " + evt.Description2 + "\n")
	sb.WriteString("TypeID: " + evt.TypeID + "\n")
	sb.WriteString("Type: " + evt.Type + "\n")
	sb.WriteString("Date: " + evt.Date.Format("2006-01-02 15:04") + "\n")
	sb.WriteString("Subscribers: " + strconv.Itoa(len(evt.Subscribers)) + "\n")

	subBtn := telebot.InlineButton{
		Unique: "subscribe_event_" + eventID,
		Text:   "✅ Подписаться",
	}

	// Отправляем с HTML-парсингом
	return c.Reply(
		sb.String(),
		&telebot.SendOptions{
			ParseMode: telebot.ModeHTML,
		},
		&telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.InlineButton{{subBtn}},
		},
	)
}
