package event_status

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func ArchiveEventService(userID int64, localID int) models.Response {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Не удалось найти событие.",
			UserDetails:     "Проверьте ID, возможно оно неверное.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	event.Status = models.StatusArchived
	if err := r_event.ReplaceEvent(event); err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Ошибка архивации события.",
			UserDetails:     "Не удалось установить статус 'архив'.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Событие успешно архивировано.",
		UserDetails:    "Теперь оно скрыто и не активно.",
		MinVisibleMode: config.ProdMode,
		VisibleToUser:  true,
	}
}
