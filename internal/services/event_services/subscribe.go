package event_services

import (
	"errors"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func Subscribe(event *models.Event, userID int64) error {
	for _, id := range event.Subscribers {
		if id == userID {
			return errors.New("вы уже подписаны")
		}
	}
	event.Subscribers = append(event.Subscribers, userID)
	return r_event.UpdateEvent(event)
}

func Unsubscribe(event *models.Event, userID int64) error {
	newArr := make([]int64, 0, len(event.Subscribers))
	for _, id := range event.Subscribers {
		if id != userID {
			newArr = append(newArr, id)
		}
	}
	event.Subscribers = newArr
	return r_event.UpdateEvent(event)
}
