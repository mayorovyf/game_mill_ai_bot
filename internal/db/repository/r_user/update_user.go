// internal/db/repository/r_user/update_user.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// обновляем данные пользователя
func UpdateUser(user *models.User) error {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// создаём фильтр
	filter := bson.M{
		"id": user.ID,
	}

	// меняем данные
	_, err := collection.ReplaceOne(ctx, filter, user)

	return err
}
