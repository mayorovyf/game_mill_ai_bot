package handlers

import (
	"game_mill_ai_bot/internal/db/repository"
	"gopkg.in/telebot.v3"
	"strconv"
)

func TeamInfoHandler(c telebot.Context) error {
	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Команду нужно вызывать в топике")
	}

	chatID := strconv.FormatInt(c.Chat().ID, 10)
	threadID := strconv.FormatInt(int64(message.ThreadID), 10)

	team, err := repository.GetTeamById(chatID, threadID)
	if err != nil {
		return c.Reply("Ошибка при получении команды: " + err.Error())
	}
	if team == nil {
		return c.Reply("Команда не найдена")
	}

	// Получаем список пользователей
	users, err := repository.GetUsersByIds(team.Members)
	if err != nil {
		return c.Reply("Ошибка при получении участников: " + err.Error())
	}

	// Формируем список участников с @username
	memberList := ""
	for _, user := range users {
		if user.Username != "" {
			memberList += "@" + user.Username + "\n"
		} else {
			memberList += "(без username)\n"
		}
	}

	// Собираем финальный текст
	text := "[" + strconv.Itoa(team.Lvl) + "] " + team.Name + " (" + team.Id + ")\n\n" +
		"Участники (" + strconv.Itoa(len(users)) + "):\n" + memberList

	return c.Reply(text)
}
