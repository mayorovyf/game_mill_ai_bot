// internal/db/repository/r_user/add_user.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// добавляем пользователя
func CreateUser(user models.User) error {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// добавляем пользователя
	_, err := collection.InsertOne(ctx, user)

	return err
}
