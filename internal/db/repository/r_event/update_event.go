// internal/db/repository/r_event/update_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// заменяем старое событие на новое с сохранением id
func UpdateEvent(event *models.Event) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("events")

	filter := bson.M{"global_id": event.GlobalID}

	_, err := collection.ReplaceOne(ctx, filter, event)

	return err
}
