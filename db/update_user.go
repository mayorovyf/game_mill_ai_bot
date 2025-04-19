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
	update := bson.M{"$set": user}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
