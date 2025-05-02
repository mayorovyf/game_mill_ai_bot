// internal/db/repository/r_event/find_events_by_user.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ищем события пользователя
func FindEventsByUser(userID int64) ([]*models.Event, error) {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{
		"author_id": userID,
	}

	// ищем события
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// формируем массив
	var events []*models.Event
	for cursor.Next(ctx) {
		var event models.Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}

	return events, nil
}
