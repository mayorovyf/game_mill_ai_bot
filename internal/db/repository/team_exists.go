package repository

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"time"
)

// TeamExists проверяет, существует ли команда с таким ID
func TeamExists(id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := db.DB.Collection("teams")
	filter := map[string]interface{}{"id": id}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
