// internal/db/repository/r_team.go
package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// получаем пользователя по его username
func GetUserByUsername(username string) (*models.User, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{
		"username": username,
	}

	// получаем коллекцию
	collection := db.DB.Collection("users")

	// получаем пользователя
	var user models.User
	err := collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // пользователь не найден
		}
		return nil, err
	}

	return &user, nil
}
