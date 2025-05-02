package draft

import (
	"fmt"
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"math/rand"
)

func CreateDraft(userID, chatID int64, topicID *int64) models.Response {
	// генерируем новые id
	localID := rand.Intn(900) + 100
	globalID := rand.Int63n(90000000) + 10000000

	// создаём черновик
	event := &models.Event{
		GlobalID:    globalID,
		LocalID:     localID,
		AuthorID:    userID,
		Status:      models.StatusDraft,
		ChatID:      chatID,
		TopicID:     topicID,
		Subscribers: []int64{userID},
	}

	// добавляем в бд
	err := r_event.AddEvent(event)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Ошибка создания события.",
			UserDetails:     "Не удалось сохранить черновик в базу данных.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    fmt.Sprintf("Создан черновик события с ID: %d", event.LocalID),
		UserDetails:    fmt.Sprintf("Для заполнения используйте /set <поле> %d <значение>", event.LocalID),
		MinVisibleMode: config.ProdMode,
		VisibleToUser:  true,
	}
}
