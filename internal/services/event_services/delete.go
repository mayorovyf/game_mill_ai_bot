package event_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
)

func DeleteEvent(userID int64, localID int) models.Response {
	err := r_event.DeleteEventByLocalID(userID, localID)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Не удалось удалить событие.",
			UserDetails:     "Возможно, оно не существует или уже удалено.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Событие успешно удалено.",
		UserDetails:    "Удаление прошло без ошибок.",
		MinVisibleMode: config.ProdMode,
		VisibleToUser:  true,
	}
}
