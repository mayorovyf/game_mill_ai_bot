package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// получаем чат по ID
func FindChat(chatID int64) (*models.Chat, error) {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("chats")

	// создаём фильтр
	filter := bson.M{"id": chatID}

	var chat models.Chat
	err := collection.FindOne(ctx, filter).Decode(&chat)

	return &chat, err
}
