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
		}
	}

	if exists {
		return models.Response{
			Level:          models.LevelInfo,
			Description:    "Вы уже зарегистрированы.",
			UserDetails:    "Успешно",
			MinVisibleMode: config.TestMode,
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
		}
	}

	return models.Response{
		Level:          models.LevelInfo,
		Description:    "Добро пожаловать!",
		UserDetails:    "Пользователь успешно зарегистрирован.",
		MinVisibleMode: config.ProdMode,
	}
}
