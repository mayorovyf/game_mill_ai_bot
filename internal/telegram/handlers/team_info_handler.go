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

	text := "Название: " + team.Name + "\n" +
		"ID: " + team.Id + "\n" +
		"Участники: " + strconv.Itoa(len(team.Members)) + "\n" +
		"Уровень: " + strconv.Itoa(team.Lvl)

	return c.Reply(text)
}
