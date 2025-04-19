package db

import (
	"context"
	"game_mill_ai_bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func UserBalance(userId int) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil // пользователь не найден
		}
		return 0, err // другая ошибка
	}
	return user.Cloudlets, nil
}
