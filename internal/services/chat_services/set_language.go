package chat_services

import (
	"errors"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
)

func SetChatLanguage(chatID int64, lang models.Language) error {
	if lang != models.LangRU && lang != models.LangEN {
		return errors.New("неподдерживаемый язык")
	}

	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return err
	}

	chat.Language = lang
	return r_chat.ReplaceChat(chat)
}
