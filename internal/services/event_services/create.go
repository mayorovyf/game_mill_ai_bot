package event_services

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func CreateDraft(userID, chatID int64, topicID *int64) (*models.Event, error) {
	return r_event.CreateDraft(userID, chatID, topicID)
}
