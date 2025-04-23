package r_user

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// UserExists проверяет, существует ли пользователь с указанным ID
func UserExists(userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("users")

	err := collection.FindOne(ctx, bson.M{"id": userID}).Err()
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
