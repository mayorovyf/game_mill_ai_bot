package user_services

import (
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
			Code:          500,
			Level:         models.LevelError,
			Description:   "Ошибка при проверке существования пользователя: " + err.Error(),
			VisibleToUser: true,
		}
	}

	if exists {
		return models.Response{
			Code:          200,
			Level:         models.LevelInfo,
			Description:   "Пользователь уже существует.",
			VisibleToUser: true,
		}
	}

	err = r_user.AddUser(user)
	if err != nil {
		return models.Response{
			Code:          500,
			Level:         models.LevelError,
			Description:   "Ошибка при регистрации пользователя: " + err.Error(),
			VisibleToUser: true,
		}
	}

	return models.Response{
		Code:          201,
		Level:         models.LevelInfo,
		Description:   "Пользователь успешно зарегистрирован.",
		VisibleToUser: true,
	}
}
