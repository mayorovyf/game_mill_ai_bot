package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetTeamById(chatId, threadId string) (*models.Team, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := db.DB.Collection("teams")
	filter := map[string]interface{}{
		"id":     threadId,
		"chatId": chatId,
	}

	var team models.Team
	err := collection.FindOne(ctx, filter).Decode(&team)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &team, nil
}
