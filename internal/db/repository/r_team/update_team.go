// internal/db/repository/r_team/update_team.go
package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// обновляем данные команды
func UpdateTeam(team models.Team) error {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("teams")

	// создаём фильтр
	filter := bson.M{
		"chatId": team.ChatId,
		"id":     team.Id,
	}

	// меняем данные
	_, err := collection.ReplaceOne(ctx, filter, team)

	return err
}
