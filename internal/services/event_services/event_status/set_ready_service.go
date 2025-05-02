package event_status

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func SetReadyService(userID int64, localID int) models.Response {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Не удалось найти событие.",
			UserDetails:     "Возможно, вы указали неправильный ID.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	event.Status = models.StatusReady
	if err := r_event.ReplaceEvent(event); err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Ошибка обновления события.",
			UserDetails:     "Не удалось установить статус 'активно'.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Событие опубликовано и теперь активно.",
		UserDetails:    "Вы можете поделиться ID с участниками.",
		MinVisibleMode: config.ProdMode,
		VisibleToUser:  true,
	}
}
