package r_chat

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// удаляем чат по ID
func DeleteChat(chatID int64) error {

	// ограничиваем запрос в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("chats")

	// создаём фильтр
	filter := bson.M{"id": chatID}

	// удаляем документ
	_, err := collection.DeleteOne(ctx, filter)

	return err
}
