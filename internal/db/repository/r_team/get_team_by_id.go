// internal/db/repository/r_team/get_team_by_id.go
package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// получаем команду по id
func GetTeamById(chatId, threadId string) (*models.Team, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("teams")

	// создаём фильтр
	filter := map[string]interface{}{
		"id":     threadId,
		"chatId": chatId,
	}

	// получаем команду
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
