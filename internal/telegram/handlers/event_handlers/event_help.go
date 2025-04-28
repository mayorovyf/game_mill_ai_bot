package event_handlers

import (
	"gopkg.in/telebot.v3"
	"strings"
)

func EventHelpHandler(c telebot.Context) error {
	helpText := strings.TrimSpace(`
<b>Команды событий:</b>

/newevent — создать новое событие
/set <code>id</code> <code>поле</code> <code>значение</code> — изменить поле события
/events — список ваших событий
/showevent <code>id</code> — подробности события
/delete <code>id</code> — удалить событие

/subscribe <code>id</code> — подписаться на напоминания
/unsubscribe <code>id</code> — отписаться от напоминаний

/sethelp — подсказка по полям для /set
/ready <code>id</code> — сделать событие активным (готовым к публикации)

/eventhelp — эта справка
`)
	return c.Reply(helpText, &telebot.SendOptions{ParseMode: telebot.ModeHTML})
}
