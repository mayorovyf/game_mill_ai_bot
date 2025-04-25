// internal/db/repository/r_event/create_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// добавляем событие в бд
func AddEvent(event *models.Event) error {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// вставляем элемент в коллекцию
	_, err := collection.InsertOne(ctx, event)

	return err
}
