package event_services

import (
	"fmt"
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func ListUserEventsInChat(userID, chatID int64) models.Response {
	events, err := r_event.FindEventsByUser(userID)
	if err != nil {
		return utils.Error("Ошибка получения событий.", "Не удалось получить ваши события.", err)
	}

	// фильтруем по chatID
	filtered := make([]*models.Event, 0)
	for _, e := range events {
		if e.ChatID == chatID {
			filtered = append(filtered, e)
		}
	}

	if len(filtered) == 0 {
		return models.Response{
			Level:          models.LevelInfo,
			Description:    "В этом чате у вас нет событий.",
			UserDetails:    "Список событий пуст.",
			VisibleToUser:  true,
			MinVisibleMode: config.ProdMode,
		}
	}

	out := ""
	for _, ev := range filtered {
		out += fmt.Sprintf("ID: %d [%s] %s\n", ev.LocalID, ev.Status, ev.Title)
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    out,
		UserDetails:    "События отфильтрованы по чату.",
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
