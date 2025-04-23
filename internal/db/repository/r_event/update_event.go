// internal/db/repository/r_event/update_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// заменяет старое событие на новое с сохранением id
func UpdateEvent(event *models.Event) error {

	// таймаут 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// замена события в бд
	_, err := db.DB.Collection("events").ReplaceOne(ctx, bson.M{"id": event.ID}, event)

	return err
}
