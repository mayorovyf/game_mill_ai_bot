// internal/db/repository/r_event/create_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// добавление события в бд
func AddEvent(event *models.Event) error {

	// таймаут 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// вставляем элемент в коллекцию
	_, err := db.DB.Collection("events").InsertOne(ctx, event)

	return err
}
