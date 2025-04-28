package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// Найти события cо статусом ready, которые начнутся в ближайшие windowMinutes минут
func FindUpcomingEvents(windowMinutes int) ([]*models.Event, error) {
	now := time.Now().UTC()
	from := now
	to := now.Add(time.Duration(windowMinutes) * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection("events")
	filter := bson.M{
		"start_time": bson.M{
			"$gte": from,
			"$lt":  to,
		},
		"status":          "ready",
		"reminder_mins.0": bson.M{"$exists": true}, // есть хотя бы одно напоминание
	}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

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
