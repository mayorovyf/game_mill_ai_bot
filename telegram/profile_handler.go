package telegram

import (
	"game_mill_ai_bot/db"
	"gopkg.in/telebot.v3"
	"strconv"
)

func ProfileHandler(c telebot.Context) error {
	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Send("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Send("Пожалуйста, используй команду в топике")
	}

	userId := strconv.FormatInt(c.Sender().ID, 10)

	exist, _ := db.UserExists(userId)
	if !exist {
		return nil
	}

	user, _ := db.GetUserById(userId)

	reply := "id: " + userId + "\n" + "admin lvl: " + strconv.FormatInt(int64(user.Adminlvl), 10) + "\n" + "облачка: " + strconv.FormatInt(int64(user.Cloudlets), 10)

	return c.Reply(reply)
}
