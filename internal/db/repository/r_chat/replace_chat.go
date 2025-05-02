package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// заменяем старый чат новым с сохранением id
func ReplaceChat(chat *models.Chat) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("chats")

	// создаём фильтр
	filter := bson.M{"id": chat.ID}

	// заменяем чат
	_, err := collection.ReplaceOne(ctx, filter, chat)

	return err
}
