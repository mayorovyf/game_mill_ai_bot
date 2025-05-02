// internal/db/repository/r_event/replace_event.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// заменяем старое событие на новое с сохранением id
func ReplaceEvent(event *models.Event) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{"global_id": event.GlobalID}

	// заменяем событие
	_, err := collection.ReplaceOne(ctx, filter, event)

	return err
}
