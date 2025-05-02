// internal/db/repository/r_event/update_reminders_sent.go
package r_event

import (
	"context"
	"time"

	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

// обновляем значения уже отправленных минут
func UpdateEventRemindersSent(globalID int64, remindersSent map[string]bool) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("events")

	// создаём фильтр
	filter := bson.M{"global_id": globalID}

	// задаём действие
	update := bson.M{
		"$set": bson.M{
			"reminders_sent": remindersSent,
		},
	}

	// обновляем событие
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
