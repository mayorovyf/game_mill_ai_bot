package telegram

import (
	"game_mill_ai_bot/db"
	"gopkg.in/telebot.v3"
	"strconv"
)

func AddCloudletsHandler(c telebot.Context) error {
	message := c.Message()

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}

	if message.ReplyTo == nil {
		return c.Reply("Пожалуйста, ответьте на сообщение пользователя, которому хотите изменить облачки")
	}

	// ID исполнителя команды (тот, кто добавляет/удаляет)
	adminId := strconv.FormatInt(c.Sender().ID, 10)

	// Проверка прав администратора
	adminLvl, err := db.UserPermissionLevel(adminId)
	if err != nil {
		return c.Reply("Ошибка при проверке уровня доступа")
	}
	if adminLvl != 99 {
		return c.Reply("У вас недостаточно прав для выполнения этой команды")
	}

	// ID цели
	targetUserId := strconv.FormatInt(message.ReplyTo.Sender.ID, 10)

	args := c.Args()
	if len(args) != 1 {
		return c.Reply("Укажите количество облачков, например: /add 10 или /add -5")
	}

	amount, err := strconv.Atoi(args[0])
	if err != nil {
		return c.Reply("Неверное значение облачков")
	}
	if amount == 0 {
		return nil
	}

	// Проверка существования пользователя
	exist, err := db.UserExists(targetUserId)
	if err != nil {
		return c.Reply("Ошибка при проверке пользователя")
	}
	if !exist {
		return c.Reply("Тот кому вы хотите начислить облачка не зарегистрирован в боте")
	}

	// Получаем пользователя
	user, err := db.GetUserById(targetUserId)
	if err != nil {
		return c.Reply("Ошибка при получении пользователя")
	}

	// Расчёт нового баланса
	newBalance := user.Cloudlets + amount
	if newBalance < 0 {
		newBalance = 0
	}
	user.Cloudlets = newBalance

	// Обновляем в базе
	err = db.UpdateUser(user)
	if err != nil {
		return c.Reply("Ошибка при обновлении пользователя")
	}

	username := message.ReplyTo.Sender.Username
	if username == "" {
		username = "пользователю"
	} else {
		username = "@" + username
	}

	action := "Начислено"
	if amount < 0 {
		action = "Списано"
		amount = -amount
	}

	return c.Reply(action + " " + strconv.Itoa(amount) + " 🌥 " + username + "\nБаланс: " + strconv.Itoa(user.Cloudlets))
}
