package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetEventById(eventID string) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var event models.Event
	err := db.DB.Collection("events").FindOne(ctx, bson.M{"id": eventID}).Decode(&event)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	return &event, err
}
