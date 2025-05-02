package event_services

import (
	"fmt"
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
	"strings"
)

func GetEventCard1(userID int64, localID int) models.Response {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil || event == nil {
		return utils.Error("Событие не найдено.", "Проверьте ID и попробуйте снова.", err)
	}

	reminders := strings.Trim(strings.Replace(fmt.Sprint(event.ReminderMins), " ", ", ", -1), "[]")

	description := fmt.Sprintf(
		`📌 <b>%s</b>
🕒 %s UTC
📝 %s
🔔 Напоминания: %s
👥 Подписчики: %d`,
		event.Title,
		event.StartTime.Format("2006-01-02 15:04"),
		event.Description,
		reminders,
		len(event.Subscribers),
	)

	return models.Response{
		Level:          models.LevelInfo,
		Description:    description,
		UserDetails:    fmt.Sprintf("Информация о событии #%d", event.LocalID),
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
