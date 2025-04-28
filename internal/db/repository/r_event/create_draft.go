package r_event

import (
	"game_mill_ai_bot/internal/models"
	"math/rand"
)

func CreateDraft(userID, chatID int64, topicID *int64) (*models.Event, error) {
	localID := rand.Intn(900) + 100
	globalID := rand.Int63n(90000000) + 10000000
	event := &models.Event{
		GlobalID:    globalID,
		LocalID:     localID,
		AuthorID:    userID,
		Status:      models.StatusDraft,
		ChatID:      chatID,
		TopicID:     topicID,
		Subscribers: []int64{userID},
	}
	err := AddEvent(event)
	return event, err
}
