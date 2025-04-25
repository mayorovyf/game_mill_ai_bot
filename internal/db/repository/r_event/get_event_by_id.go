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

// получаем событие по id
func GetEventById(eventID string) (*models.Event, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// создаём фильтр
	filter := bson.M{
		"id": eventID,
	}

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// поиск события в бд
	var event models.Event
	err := collection.FindOne(ctx, filter).Decode(&event)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &event, err
}
