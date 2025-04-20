package repository

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateTeam(team models.Team) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := db.DB.Collection("teams")

	filter := bson.M{
		"chatId": team.ChatId,
		"id":     team.Id,
	}

	_, err := collection.ReplaceOne(ctx, filter, team)
	return err
}
