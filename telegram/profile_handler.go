package telegram

import (
	"game_mill_ai_bot/db"
	"gopkg.in/telebot.v3"
	"strconv"
)

func ProfileHandler(c telebot.Context) error {
	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}

	userId := strconv.FormatInt(c.Sender().ID, 10)

	exist, _ := db.UserExists(userId)
	if !exist {
		return nil
	}

	user, _ := db.GetUserById(userId)

	reply := "Профиль пользователя @" + c.Sender().Username + "\n" +
		"id: " + userId + "\n" +
		"admin lvl: " + strconv.Itoa(user.Adminlvl) + "\n" +
		"облачка: " + strconv.Itoa(user.Cloudlets)

	return c.Reply(reply)
}
