package event_services

import (
	"fmt"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils"
)

func SubscribeToEvent(userID int64, localID int) models.Response {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil || event == nil {
		return utils.Error("Событие не найдено.", "Проверьте ID и повторите попытку.", err)
	}

	// уже подписан?
	for _, id := range event.Subscribers {
		if id == userID {
			return utils.Warn("Вы уже подписаны на это событие.", "Ничего не изменилось.")
		}
	}

	event.Subscribers = append(event.Subscribers, userID)

	if err := r_event.ReplaceEvent(event); err != nil {
		return utils.Error("Не удалось сохранить подписку.", "Попробуйте позже.", err)
	}

	return utils.Info(
		fmt.Sprintf("Вы подписались на событие: %s", event.Title),
		fmt.Sprintf("ID: %d | Подписчиков: %d", event.LocalID, len(event.Subscribers)),
	)
}

func UnsubscribeFromEvent(userID int64, localID int) models.Response {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil || event == nil {
		return utils.Error("Событие не найдено.", "Проверьте ID и повторите попытку.", err)
	}

	// отписка
	filtered := make([]int64, 0, len(event.Subscribers))
	found := false
	for _, id := range event.Subscribers {
		if id == userID {
			found = true
			continue
		}
		filtered = append(filtered, id)
	}
	if !found {
		return utils.Warn("Вы не подписаны на это событие.", "Отписка не требуется.")
	}

	event.Subscribers = filtered

	if err := r_event.ReplaceEvent(event); err != nil {
		return utils.Error("Не удалось сохранить изменения.", "Попробуйте позже.", err)
	}

	return utils.Info(
		fmt.Sprintf("Вы отписались от события: %s", event.Title),
		fmt.Sprintf("ID: %d | Осталось подписчиков: %d", event.LocalID, len(event.Subscribers)),
	)
}
