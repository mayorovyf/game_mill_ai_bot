// internal/db/repository/r_event/get_event_by_id.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// получение события по id
func GetEventById(eventID string) (*models.Event, error) {

	// таймаут 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// поиск события в бд
	var event models.Event
	err := db.DB.Collection("events").FindOne(ctx, bson.M{"id": eventID}).Decode(&event)

	// обработка случая когда событие не найдено
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &event, err
}
