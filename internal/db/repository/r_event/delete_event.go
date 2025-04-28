package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func DeleteEvent(userID int64, localID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection("events")
	filter := bson.M{"author_id": userID, "local_id": localID}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}
