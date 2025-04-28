package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func FindEventByLocalID(userID int64, localID int) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection("events")
	filter := bson.M{"author_id": userID, "local_id": localID}
	var event models.Event
	err := collection.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
