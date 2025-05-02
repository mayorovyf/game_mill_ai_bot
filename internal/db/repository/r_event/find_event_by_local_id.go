// internal/db/repository/r_event/find_event_by_local_id.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// получаем событие по локальному id у пользователя
func FindEventByLocalID(userID int64, localID int) (*models.Event, error) {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{
		"author_id": userID,
		"local_id":  localID,
	}

	// получаем событие
	var event models.Event
	err := collection.FindOne(ctx, filter).Decode(&event)

	if err != nil {
		return nil, err
	}
	return &event, nil
}
