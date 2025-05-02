package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// проверяем, существует ли чат по id
func ChatExists(chatID int64) (bool, error) {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("chats")

	count, err := collection.CountDocuments(ctx, bson.M{"id": chatID})

	return count > 0, err
}
