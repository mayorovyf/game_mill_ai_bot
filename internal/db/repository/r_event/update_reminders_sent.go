package r_event

import (
	"context"
	"time"

	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateEventRemindersSent(globalID int64, remindersSent map[string]bool) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := db.DB.Collection("events")
	filter := bson.M{"global_id": globalID}
	update := bson.M{"$set": bson.M{"reminders_sent": remindersSent}}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
