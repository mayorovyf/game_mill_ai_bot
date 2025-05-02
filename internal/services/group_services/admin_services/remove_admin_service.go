package admin_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func RemoveAdmin(chatID, userID int64) models.Response {
	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return utils.Error("Ошибка получения чата.", "Не удалось найти чат.", err)
	}

	found := false
	filtered := make([]models.GroupAdmin, 0)
	for _, admin := range chat.Admins {
		if admin.UserID == userID {
			found = true
			continue
		}
		filtered = append(filtered, admin)
	}
	if !found {
		return models.Response{
			Level:          models.LevelWarn,
			Description:    "Админ не найден.",
			UserDetails:    "Такого админа нет в чате.",
			VisibleToUser:  true,
			MinVisibleMode: config.ProdMode,
		}
	}

	chat.Admins = filtered
	if err := r_chat.ReplaceChat(chat); err != nil {
		return utils.Error("Ошибка удаления админа.", "Не удалось сохранить изменения.", err)
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Админ успешно удалён.",
		UserDetails:    "Права пользователя сняты.",
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
