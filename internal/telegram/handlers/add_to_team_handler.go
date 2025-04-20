package handlers

import (
	"game_mill_ai_bot/internal/db/repository"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func AddToTeamHandler(c telebot.Context) error {
	message := c.Message()

	// Проверка: только в супергруппах и в топиках
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}

	args := c.Args()
	if len(args) != 1 || !strings.HasPrefix(args[0], "@") {
		return c.Reply("Укажите юзернейм, например: /add_to_team @username")
	}
	username := strings.TrimPrefix(args[0], "@")

	// Получение пользователя по username
	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return c.Reply("Ошибка при поиске пользователя: " + err.Error())
	}
	if user == nil {
		return c.Reply("Пользователь с юзернеймом @" + username + " не зарегистрирован в боте")
	}

	// Получение команды
	chatID := strconv.FormatInt(c.Chat().ID, 10)
	threadID := strconv.FormatInt(int64(message.ThreadID), 10)
	team, err := repository.GetTeamById(chatID, threadID)
	if err != nil || team == nil {
		return c.Reply("Команда не найдена")
	}

	// Проверка: уже в команде?
	for _, member := range team.Members {
		if member == user.ID {
			return c.Reply("Пользователь уже в команде")
		}
	}

	// Добавляем в команду
	team.Members = append(team.Members, user.ID)

	err = repository.UpdateTeam(*team)
	if err != nil {
		return c.Reply("Ошибка при добавлении пользователя в команду: " + err.Error())
	}

	return c.Reply("Пользователь @" + username + " успешно добавлен в команду")
}
