package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// получаем список всех чатов
func ListChats() ([]*models.Chat, error) {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("chats")

	// получаем все документы
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// декодируем все чаты
	var chats []*models.Chat
	for cursor.Next(ctx) {
		var chat models.Chat
		if err := cursor.Decode(&chat); err == nil {
			chats = append(chats, &chat)
		}
	}

	return chats, cursor.Err()
}
