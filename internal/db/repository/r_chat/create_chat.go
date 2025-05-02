package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// создаём новый чат
func CreateChat(chat *models.Chat) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("chats")

	// вставляем чат
	_, err := collection.InsertOne(ctx, chat)

	return err
}
