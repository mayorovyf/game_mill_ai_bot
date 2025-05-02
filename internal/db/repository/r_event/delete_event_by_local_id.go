// internal/db/repository/r_event/delete_event_by_local_id.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// удаляем событие по локальному id у пользователя
func DeleteEventByLocalID(userID int64, localID int) error {

	// ограничиваем запрос в 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{
		"author_id": userID,
		"local_id":  localID,
	}

	// удаляем событие
	_, err := collection.DeleteOne(ctx, filter)

	return err
}
