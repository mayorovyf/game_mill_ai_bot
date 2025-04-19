package db

import (
	"context"
	"game_mill_ai_bot/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTeam(team models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Collection("teams")

	filter := bson.M{
		"chatId": team.ChatId,
		"id":     team.Id,
	}

	_, err := collection.ReplaceOne(ctx, filter, team)
	return err
}
