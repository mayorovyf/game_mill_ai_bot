package chat_services

import (
	"game_mill_ai_bot/internal/db/repository/r_chat"
)

func SetChatTimezone(chatID int64, tz int) error {
	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return err
	}

	chat.Timezone = tz
	return r_chat.ReplaceChat(chat)
}
