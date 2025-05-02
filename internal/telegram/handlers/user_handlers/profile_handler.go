package user_handlers

import (
	"game_mill_ai_bot/internal/db/repository/r_user"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"gopkg.in/telebot.v3"
	"strconv"
)

func ProfileHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}

	userId := strconv.FormatInt(c.Sender().ID, 10)

	exist, _ := r_user.UserExists(userId)
	if !exist {
		c.Reply("Вы не зарегистрированы")
		return nil
	}

	user, _ := r_user.GetUserById(userId)

	reply := "Профиль пользователя @" + c.Sender().Username + "\n" +
		"id: " + userId + "\n" +
		"admin lvl: " + strconv.Itoa(user.Adminlvl) + "\n" +
		"облачка: " + strconv.Itoa(user.Cloudlets)

	return c.Reply(reply)
}
