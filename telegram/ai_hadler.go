package telegram

import (
	"game_mill_ai_bot/ai"
	"game_mill_ai_bot/config"
	"game_mill_ai_bot/db"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
)

func AiHendler(c telebot.Context) error {
	message := c.Message()
	prompt := strings.TrimSpace(message.Payload)

	if c.Chat().Type != telebot.ChatSuperGroup {
		return c.Reply("Бот работает только в супергруппах")
	}
	if message.ThreadID == 0 {
		return c.Reply("Пожалуйста, используй команду в топике")
	}
	if prompt == "" {
		return c.Reply("Пожалуйста, укажи запрос после команды, например:\n`/ai Что такое черная дыра?`", &telebot.SendOptions{ParseMode: telebot.ModeMarkdown})
	}

	// Получаем ID пользователя
	userId := strconv.FormatInt(c.Sender().ID, 10)

	// Проверка существования
	exists, err := db.UserExists(userId)
	if err != nil {
		return c.Reply("Ошибка при проверке пользователя")
	}
	if !exists {
		return c.Reply("Вы не зарегистрированы в системе. Попросите администратора добавить вас.")
	}

	// Получение пользователя
	user, err := db.GetUserById(userId)
	if err != nil {
		return c.Reply("Ошибка при получении пользователя")
	}

	// Проверка баланса
	if user.Cloudlets <= 0 {
		return c.Reply("У вас недостаточно облачков для использования ИИ 😔")
	}

	// Отправка предварительного сообщения
	sendOpts := &telebot.SendOptions{
		ThreadID:  message.ThreadID,
		ParseMode: telebot.ModeMarkdown,
	}
	c.Send("Думаю...", sendOpts)

	// Получение ответа от ИИ
	reply, err := ai.GetChatResponse(config.AiClient, prompt)
	if err != nil {
		return c.Reply("ИИ не смогла ответить")
	}

	// Списание 1 облачка
	user.Cloudlets -= 1
	err = db.UpdateUser(user)
	if err != nil {
		return c.Reply("Ошибка при обновлении баланса пользователя", sendOpts)
	}

	return c.Send(reply, sendOpts)
}
