// internal/db/repository/r_event/add_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// добавляем событие в бд
func AddEvent(event *models.Event) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// задаём коллекцию
	collection := db.DB.Collection("events")

	// добавляем элемент
	_, err := collection.InsertOne(ctx, event)

	return err
}
