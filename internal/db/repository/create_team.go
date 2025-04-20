package repository

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// CreateTeam добавляет новую команду, если её ещё нет
func CreateTeam(team models.Team) error {
	exists, err := TeamExists(team.Id)
	if err != nil {
		return err
	}
	if exists {
		return nil // команда уже существует — не добавляем
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := db.DB.Collection("teams")
	_, err = collection.InsertOne(ctx, team)
	return err
}
