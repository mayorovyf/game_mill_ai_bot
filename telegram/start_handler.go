package telegram

import (
	"game_mill_ai_bot/db"
	"game_mill_ai_bot/models"
	"github.com/sashabaranov/go-openai"
	"gopkg.in/telebot.v3"
	"strconv"
)

func StartHandler(c telebot.Context) error {
	if c.Chat().Type != telebot.ChatPrivate {
		return c.Reply("Эта команда работает только в лс")
	}

	user := models.User{
		ID:           strconv.FormatInt(c.Sender().ID, 10),
		Username:     c.Sender().Username,
		Cloudlets:    0,
		Adminlvl:     0,
		CurrentModel: openai.GPT3Dot5Turbo,
	}

	exists, err := db.UserExists(user.ID)
	if exists {
		return nil
	}

	err = db.CreateUser(user)
	if err != nil {
		return c.Reply(" Ошибка при регистрации пользователя:\n" + err.Error())
	}

	return c.Reply(" Добро пожаловать!")
}
