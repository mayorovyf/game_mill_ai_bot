package chat_services

import (
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
)

// InitChat создаёт или обновляет чат
func InitChat(chat *models.Chat) error {
	exists, err := r_chat.ChatExists(chat.ID)
	if err != nil {
		return err
	}

	if exists {
		return r_chat.ReplaceChat(chat)
	}

	return r_chat.CreateChat(chat)
}
