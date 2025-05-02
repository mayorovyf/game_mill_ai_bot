package admin_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func AssignAdmin(chatID, userID int64, title string) models.Response {
	chat, err := r_chat.FindChat(chatID)
	if err != nil {
		return utils.Error("Ошибка получения чата.", "Не удалось найти чат.", err)
	}

	for _, admin := range chat.Admins {
		if admin.UserID == userID {
			return models.Response{
				Level:          models.LevelWarn,
				Description:    "Пользователь уже админ.",
				UserDetails:    "Этот пользователь уже является админом.",
				VisibleToUser:  true,
				MinVisibleMode: config.ProdMode,
			}
		}
	}

	newAdmin := models.GroupAdmin{
		UserID: userID,
		Title:  title,
		Privileges: models.AdminPrivileges{
			CanEditEvents:     true,
			CanManageMembers:  false,
			CanArchive:        true,
			CanDeleteMessages: false,
		},
	}

	chat.Admins = append(chat.Admins, newAdmin)

	if err := r_chat.ReplaceChat(chat); err != nil {
		return utils.Error("Ошибка назначения админа.", "Не удалось сохранить изменения.", err)
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Админ успешно назначен.",
		UserDetails:    title,
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
