package admin_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func SetAdminTitle(chatID, userID int64, title string) models.Response {
	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return utils.Error("Ошибка получения чата.", "Не удалось найти чат.", err)
	}

	for i := range chat.Admins {
		if chat.Admins[i].UserID == userID {
			chat.Admins[i].Title = title
			if err := r_chat.ReplaceChat(chat); err != nil {
				return utils.Error("Ошибка изменения звания.", "Не удалось сохранить звание.", err)
			}
			return models.Response{
				Level:          models.LevelInfo,
				Description:    "Звание обновлено.",
				UserDetails:    title,
				VisibleToUser:  true,
				MinVisibleMode: config.ProdMode,
			}
		}
	}

	return models.Response{
		Level:          models.LevelWarn,
		Description:    "Админ не найден.",
		UserDetails:    "Пользователь не является админом.",
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
