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

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("teams")

	// добавляем команду
	_, err := collection.InsertOne(ctx, team)

	return err
}
