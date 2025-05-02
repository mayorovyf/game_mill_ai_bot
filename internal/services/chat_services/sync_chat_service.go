package chat_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_chat"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
	"gopkg.in/telebot.v3"
)

// SyncChat обновляет или создаёт чат в базе и возвращает структурированный ответ
func SyncChat(tbChat *telebot.Chat) models.Response {
	existing, err := r_chat.FindChat(tbChat.ID)
	if err != nil {
		// Чат не найден — пробуем создать
		newChat := &models.Chat{
			ID:       tbChat.ID,
			Type:     tbChat.Type,
			Title:    tbChat.Title,
			Username: tbChat.Username,
			Timezone: 0,
			Language: models.LangRU,
		}

		if err := r_chat.CreateChat(newChat); err != nil {
			return utils.Error(
				"Не удалось создать чат.",
				"Ошибка при первичной регистрации чата.",
				err,
			)
		}

		return models.Response{
			Level:          models.LevelInfo,
			Description:    "Создан новый чат.",
			UserDetails:    "Этот чат зарегистрирован в системе.",
			VisibleToUser:  false,
			MinVisibleMode: config.DevMode,
		}
	}

	// Проверяем изменения
	updated := false
	if existing.Type != tbChat.Type {
		existing.Type = tbChat.Type
		updated = true
	}
	if existing.Title != tbChat.Title {
		existing.Title = tbChat.Title
		updated = true
	}
	if existing.Username != tbChat.Username {
		existing.Username = tbChat.Username
		updated = true
	}

	if updated {
		if err := r_chat.ReplaceChat(existing); err != nil {
			return utils.Error(
				"Не удалось обновить данные чата.",
				"Ошибка при синхронизации данных чата.",
				err,
			)
		}

		return models.Response{
			Level:          models.LevelInfo,
			Description:    "Данные чата обновлены.",
			UserDetails:    "Информация синхронизирована.",
			VisibleToUser:  false,
			MinVisibleMode: config.DevMode,
		}
	}

	// Нет изменений
	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Данные чата актуальны.",
		UserDetails:    "Нет необходимости в обновлении.",
		VisibleToUser:  false,
		MinVisibleMode: config.DevMode,
	}
}
