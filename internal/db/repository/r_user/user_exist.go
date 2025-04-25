// internal/db/repository/r_user/user_exist.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// проверям, существует ли пользователь с таким id
func UserExists(userID string) (bool, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// создаём фильтр
	filter := bson.M{
		"id": userID,
	}

	// ищем пользователя
	err := collection.FindOne(ctx, filter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
