package event_handlers

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"gopkg.in/telebot.v3"
	"strconv"
)

func SubscribeToEvent(c telebot.Context, eventID string) error {
	evt, err := r_event.GetEventById(eventID)
	if err != nil || evt == nil {
		return c.Respond(&telebot.CallbackResponse{Text: "Событие не найдено."})
	}
	userID := strconv.FormatInt(c.Sender().ID, 10)
	for _, id := range evt.Subscribers {
		if id == userID {
			return c.Respond(&telebot.CallbackResponse{Text: "Вы уже подписаны!"})
		}
	}
	evt.Subscribers = append(evt.Subscribers, userID)
	r_event.UpdateEvent(evt)
	return c.Respond(&telebot.CallbackResponse{Text: "Вы подписались на событие!"})
}
