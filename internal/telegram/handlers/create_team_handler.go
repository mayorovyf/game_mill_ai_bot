package handlers

import (
	"game_mill_ai_bot/internal/db/repository"
	"game_mill_ai_bot/internal/models"
	"gopkg.in/telebot.v3"
	"strconv"
)

func CreateTeamHandler(c telebot.Context) error {
	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}

	adminId := strconv.FormatInt(c.Sender().ID, 10)
	adminLvl, err := repository.UserPermissionLevel(adminId)
	if err != nil {
		return c.Reply("Ошибка при проверке уровня доступа")
	}
	if adminLvl < 99 {
		return c.Reply("У вас недостаточно прав для выполнения этой команды")
	}

	team := models.Team{
		Name:    strconv.FormatInt(int64(message.ThreadID), 10), // временное имя — ID топика
		Id:      strconv.FormatInt(int64(message.ThreadID), 10), // можно использовать тот же ID как уникальный идентификатор
		ChatId:  strconv.FormatInt(c.Chat().ID, 10),             // ID супергруппы
		Members: []string{},
		Admins:  []string{},
		Lvl:     1, // стартовый уровень
	}

	err = repository.CreateTeam(team)
	if err != nil {
		return c.Reply("Ошибка при создании команды: " + err.Error())
	}

	return c.Reply("Команда успешно создана ✅")
}
