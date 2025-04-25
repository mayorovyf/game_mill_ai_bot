// internal/db/repository/r_team/create_team.go
package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// создание команды
func CreateTeam(team models.Team) error {

	// проверяем существует ли команда
	exists, err := TeamExist(team.Id, team.ChatId)

	if err != nil {
		return err
	}
	if exists {
		return nil // команда уже существует — не добавляем
	}

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// добавляем команду
	_, err = db.DB.Collection("teams").InsertOne(ctx, team)
	return err
}
