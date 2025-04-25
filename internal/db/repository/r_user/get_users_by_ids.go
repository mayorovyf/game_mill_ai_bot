// internal/db/repository/r_user/get_users_by_ids.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// получаем массив пользователей по массиву id
func GetUsersByIds(userIds []string) ([]*models.User, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// создаём фильтр
	filter := bson.M{
		"id": bson.M{"$in": userIds},
	}

	// получаем коллекцию
	collection := db.DB.Collection("users")

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// получаем пользователей
	var users []*models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
