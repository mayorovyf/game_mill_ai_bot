// internal/db/repository/r_user/user_permission_level.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// !!! Планируется удалить !!!

// получаем уровень доступа у пользователя
func UserPermissionLevel(userId string) (int, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// ищем пользователя
	var user models.User
	err := collection.FindOne(ctx, bson.M{"id": userId}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil // пользователь не найден
		}
		return 0, err
	}

	return user.Adminlvl, nil
}
