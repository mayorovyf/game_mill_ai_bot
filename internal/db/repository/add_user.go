package repository

import (
	"context"
	"game_mill_ai_bot/internal/db"
	"game_mill_ai_bot/internal/models"
	"time"
)

// CreateUser добавляет нового пользователя, если его ещё нет
func CreateUser(user models.User) error {
	exists, err := UserExists(user.ID)
	if err != nil {
		return err
	}
	if exists {
		return nil // уже есть — не добавляем
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	collection := db.DB.Collection("users")
	_, err = collection.InsertOne(ctx, user)
	return err
}
