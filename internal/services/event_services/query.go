package event_services

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func ListUserEvents(userID int64) ([]*models.Event, error) {
	return r_event.FindEventsByUser(userID)
}

func GetEventCard(userID int64, localID int) (*models.Event, error) {
	return r_event.FindEventByLocalID(userID, localID)
}
