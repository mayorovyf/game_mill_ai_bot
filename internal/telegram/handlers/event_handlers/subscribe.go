package event_handlers

import (
	"game_mill_ai_bot/internal/services/event_services"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SubscribeHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /subscribe <id>")
	}
	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}
	ev, err := event_services.GetEventCard(c.Sender().ID, localID)
	if err != nil {
		return c.Reply("Событие не найдено.")
	}
	if err := event_services.Subscribe(ev, c.Sender().ID); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Вы подписались на событие.")
}

func UnsubscribeHandler(c telebot.Context) error {
	args := strings.Fields(c.Message().Payload)
	if len(args) < 1 {
		return c.Reply("Пример: /unsubscribe <id>")
	}
	localID, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("ID должен быть числом")
	}
	ev, err := event_services.GetEventCard(c.Sender().ID, localID)
	if err != nil {
		return c.Reply("Событие не найдено.")
	}
	if err := event_services.Unsubscribe(ev, c.Sender().ID); err != nil {
		return c.Reply("Ошибка: " + err.Error())
	}
	return c.Reply("Вы отписались от события.")
}
