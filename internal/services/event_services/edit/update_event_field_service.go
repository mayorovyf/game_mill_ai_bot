package edit

import (
	"fmt"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
	"strconv"
	"time"
)

type fieldUpdater func(event *models.Event, value string) error

var updaters = map[string]fieldUpdater{
	"title": func(e *models.Event, val string) error {
		e.Title = val
		return nil
	},
	"description": func(e *models.Event, val string) error {
		e.Description = val
		return nil
	},
	"time": func(e *models.Event, val string) error {
		t, err := time.Parse("2006-01-02 15:04", val)
		if err != nil {
			return err
		}
		e.StartTime = t
		return nil
	},
	"reminder": func(e *models.Event, val string) error {
		mins, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		e.ReminderMins = append(e.ReminderMins, mins)
		return nil
	},
	"topic": func(e *models.Event, val string) error {
		topicID, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return err
		}
		e.TopicID = &topicID
		return nil
	},
}

func UpdateEventField(userID int64, localID int, field string, value string) models.Response {
	updater, exists := updaters[field]
	if !exists {
		return utils.Warn("Неизвестное поле: "+field, "Допустимые поля: title, description, time, reminder, topic")
	}

	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return utils.Error("Не удалось найти событие.", "Проверьте правильность ID.", err)
	}

	if err := updater(event, value); err != nil {
		return utils.Error("Ошибка обновления значения.", "Проверьте формат: "+field, err)
	}

	if err := r_event.ReplaceEvent(event); err != nil {
		return utils.Error("Ошибка сохранения.", "Не удалось записать изменения в базу.", err)
	}

	return utils.Info(
		fmt.Sprintf("Поле '%s' обновлено успешно.", field),
		fmt.Sprintf("ID: %d | Значение: %s", event.LocalID, field),
	)
}
