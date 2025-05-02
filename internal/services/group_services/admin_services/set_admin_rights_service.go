package admin_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func SetAdminRights(chatID, userID int64, priv models.AdminPrivileges) models.Response {
	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return utils.Error("Ошибка получения чата.", "Не удалось найти чат.", err)
	}

	for i := range chat.Admins {
		if chat.Admins[i].UserID == userID {
			chat.Admins[i].Privileges = priv
			if err := r_chat.ReplaceChat(chat); err != nil {
				return utils.Error("Ошибка обновления прав.", "Не удалось сохранить изменения.", err)
			}
			return models.Response{
				Level:          models.LevelInfo,
				Description:    "Права успешно обновлены.",
				UserDetails:    "Привилегии сохранены.",
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
