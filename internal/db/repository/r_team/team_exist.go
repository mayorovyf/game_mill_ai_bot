// internal/db/repository/r_team/team_exist
package r_team

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// проверяем, существует ли команда с таким ID
func TeamExist(id string, chat_id string) (bool, error) {

	// ограничиваем запрос к бд в 5 сек
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// создаём фильтр
	filter := bson.M{
		"id":      id,
		"chat_id": chat_id,
	}

	// получаем коллекцию
	collection := db.DB.Collection("teams")

	// проверяем наличие
	err := collection.FindOne(ctx, filter).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
