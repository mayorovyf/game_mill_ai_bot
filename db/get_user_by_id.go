package db

import (
	"context"
	"game_mill_ai_bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetUserById(userId string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"id": userId}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // пользователь не найден
		}
		return nil, err // другая ошибка
	}
	return &user, nil
}
