package help

import (
	"gopkg.in/telebot.v3"
	"strings"
)

func EventHelpHandler(c telebot.Context) error {
	helpText := strings.TrimSpace(`
<b> Команды управления событиями:</b>

/newevent — создать новое событие
/set &lt;поле&gt; &lt;id&gt; &lt;значение&gt; — изменить поле (введите /sethelp для списка полей)
/events — список всех ваших событий
/showevent &lt;id&gt; — подробности события
/delete &lt;id&gt; — удалить событие

<b> Напоминания:</b>
/subscribe &lt;id&gt; — подписаться на событие
/unsubscribe &lt;id&gt; — отписаться от события

<b> Публикация:</b>
/ready &lt;id&gt; — сделать событие активным

<b> Справка:</b>
/sethelp — помощь по полям для /set
/eventhelp — эта справка
`)

	return c.Reply(helpText, &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	})
}
