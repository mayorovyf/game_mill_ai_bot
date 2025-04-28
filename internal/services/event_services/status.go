package event_services

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func SetReady(userID int64, localID int) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.Status = models.StatusReady
	return r_event.UpdateEvent(event)
}

func Archive(userID int64, localID int) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.Status = models.StatusArchived
	return r_event.UpdateEvent(event)
}
