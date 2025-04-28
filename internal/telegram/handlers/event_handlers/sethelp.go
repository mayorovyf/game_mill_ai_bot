package event_handlers

import (
	"gopkg.in/telebot.v3"
	"strings"
)

func SetHelpHandler(c telebot.Context) error {
	helpText := strings.TrimSpace(`
<b>Справка по <code>/set &lt;поле&gt; &lt;id&gt; &lt;значение&gt;</code>:</b>

Доступные поля:
• <b>title</b> — название события (текст)
  Пример: <code>/set title 123 День рождения</code>

• <b>description</b> — описание события (текст)
  Пример: <code>/set description 123 Встреча у меня дома</code>

• <b>time</b> — время события в формате <code>YYYY-MM-DD HH:MM UTC</code>
  Пример: <code>/set time 123 2024-07-01 18:00 UTC</code>

• <b>reminder</b> — добавить напоминание (минуты до события, целое число)
  Пример: <code>/set reminder 123 60</code>

• <b>topic</b> — ID топика (целое число, опционально для групп)
  Пример: <code>/set topic 123 55555</code>

Для удаления напоминания используйте: <code>/unset reminder &lt;id&gt; &lt;минуты&gt;</code>
  Пример: <code>/unset reminder 123 60</code>

Для просмотра событий: <code>/events</code>
Для подробностей по событию: <code>/showevent &lt;id&gt;</code>
`)
	return c.Reply(helpText, &telebot.SendOptions{ParseMode: telebot.ModeHTML})
}
