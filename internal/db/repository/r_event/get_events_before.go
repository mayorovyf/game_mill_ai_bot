// intenal/db/repository/r_event/get_events_before.go
package r_event

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// получаем все события до определённого времени
func GetEventsBefore(t time.Time) ([]models.Event, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// создаём фильтр
	filter := bson.M{
		"date": bson.M{"$lte": t},
	}

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// ищем события
	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	// закрытие курсора
	defer cursor.Close(ctx)

	// получение всех событий из курсора
	var events []models.Event
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
