// interanl/services/user_services/sreate_user_service.go
package user_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/db/repository/r_user"
	"game_mill_ai_bot/internal/models"
	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
	"strconv"
)

func CreateUser(u *telebot.User) models.Response {
	user := models.User{
		ID:           strconv.FormatInt(u.ID, 10),
		Username:     u.Username,
		Cloudlets:    0,
		Adminlvl:     0,
		CurrentModel: openai.GPT4oMini,
	}

	exists, err := r_user.UserExists(user.ID)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Ошибка регистрации.",
			UserDetails:     "Не удалось проверить, зарегистрирован ли пользователь.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.TestMode,
			VisibleToUser:   true,
		}
	}

	if exists {
		return models.Response{
			Level:          models.LevelInfo,
			Description:    "Вы уже зарегистрированы.",
			UserDetails:    "Вы уже нажимали /start в этом боте",
			MinVisibleMode: config.TestMode,
			VisibleToUser:  true,
		}
	}

	err = r_user.AddUser(user)
	if err != nil {
		return models.Response{
			Level:           models.LevelError,
			Description:     "Произошла ошибка.",
			UserDetails:     "Не удалось зарегистрировать пользователя.",
			InternalDetails: err.Error(),
			MinVisibleMode:  config.DevMode,
			VisibleToUser:   true,
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Добро пожаловать!",
		UserDetails:    "Пользователь успешно зарегистрирован.",
		MinVisibleMode: config.ProdMode,
		VisibleToUser:  true,
	}
}
