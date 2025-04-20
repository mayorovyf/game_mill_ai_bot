package handlers

import (
	"game_mill_ai_bot/internal/db/repository"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func SetTeamNameHandler(c telebot.Context) error {
	message := c.Message()
	const RequiredAdminLevel = 99

	// Проверка: только в супергруппах и в топиках
	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}

	// Парсинг нового имени
	args := c.Args()
	if len(args) == 0 {
		return c.Reply("Укажите новое название команды, например: /set_team_name Моя Команда")
	}
	newName := strings.Join(args, " ")

	// Ограничение длины названия
	if len([]rune(newName)) >= 16 {
		return c.Reply("Название слишком длинное. Максимум 15 символов." + string([]rune(newName)))
	}

	// Проверка прав (допустим, нужен 99 уровень)
	adminID := strconv.FormatInt(c.Sender().ID, 10)
	adminLvl, err := repository.UserPermissionLevel(adminID)
	if err != nil {
		return c.Reply("Ошибка при проверке уровня доступа")
	}
	if adminLvl < RequiredAdminLevel {
		return c.Reply("У вас недостаточно прав для изменения названия")
	}

	// Получаем текущую команду
	chatID := strconv.FormatInt(c.Chat().ID, 10)
	threadID := strconv.FormatInt(int64(message.ThreadID), 10)

	team, err := repository.GetTeamById(chatID, threadID)
	if err != nil {
		return c.Reply("Ошибка при получении команды")
	}
	if team == nil {
		return c.Reply("Команда не найдена")
	}

	// Обновляем название и сохраняем команду
	team.Name = newName

	err = repository.UpdateTeam(*team)
	if err != nil {
		return c.Reply("Ошибка при обновлении команды: " + err.Error())
	}

	return c.Reply("Название команды обновлено: " + newName)
}
