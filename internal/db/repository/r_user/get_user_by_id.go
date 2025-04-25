// interanal/db/repository/r_user/get_user_by_id.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// получаем пользователя по id
func GetUserById(userId string) (*models.User, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// создаём фильтр
	filter := bson.M{
		"id": userId,
	}

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// ищем пользователя
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
