// internal/db/repository/r_event/find_upcoming.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// НИщем события cо статусом ready, которые начнутся в ближайшие windowMinutes минут
func FindUpcomingEvents(windowMinutes int) ([]*models.Event, error) {

	// задаём временные рамки для поиска
	now := time.Now().UTC()
	from := now
	to := now.Add(time.Duration(windowMinutes) * time.Minute)

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{
		"start_time": bson.M{
			"$gte": from,
			"$lt":  to,
		},
		"event_status":    "ready",
		"reminder_mins.0": bson.M{"$exists": true}, // есть хотя бы одно напоминание
	}

	// ищем события
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// формируем итоговый массив
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
