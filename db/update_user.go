package db

import (
	"context"
	"game_mill_ai_bot/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func UpdateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Collection("users")

	filter := bson.M{"id": user.ID}

	_, err := collection.ReplaceOne(ctx, filter, user)
	return err
}
