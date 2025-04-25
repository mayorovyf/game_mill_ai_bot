// internal/db/repository/r_team/team_exist
package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"time"
)

// проверяем, существует ли команда с таким ID
func TeamExist(id string, chat_id string) (bool, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// получаем коллекцию
	collection := db.DB.Collection("teams")

	// создаём фильтр
	filter := map[string]interface{}{
		"id":      id,
		"chat_id": chat_id,
	}

	// проверяем наличие
	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
