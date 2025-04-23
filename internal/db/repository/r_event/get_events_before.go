package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func GetEventsBefore(t time.Time) ([]models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := db.DB.Collection("events").Find(ctx, bson.M{"date": bson.M{"$lte": t}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}
