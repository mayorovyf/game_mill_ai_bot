package help

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"gopkg.in/telebot.v3"
	"strings"
)

func SetHelpHandler(c telebot.Context) error {

	response := chat_services.SyncChat(c.Chat())
	if response.Level == models.LevelError {
		return c.Reply(response.UserDetails)
	}

	helpText := strings.TrimSpace(`
<b>️ Справка по команде /set:</b>

Формат:
<code>/set &lt;поле&gt; &lt;id&gt; &lt;значение&gt;</code>

<b>Доступные поля:</b>
• <b>title</b> — название события  
  <i>Пример:</i> <code>/set</code> <code>title</code> <code>123 День рождения</code>

• <b>description</b> — описание события  
  <i>Пример:</i> <code>/set</code> <code>description</code> <code>123 Встреча у меня дома</code>

• <b>time</b> — время события в формате <code>YYYY-MM-DD HH:MM</code> (UTC)  
  <i>Пример:</i> <code>/set</code> <code>time</code> <code>123 2024-07-01 18:00</code>

• <b>reminder</b> — добавить напоминание (в минутах до события)  
  <i>Пример:</i> <code>/set</code> <code>reminder</code> <code>123 60</code>

• <b>topic</b> — ID топика (для групп/форумов)  
  <i>Пример:</i> <code>/set</code> <code>topic</code> <code>123 55555</code>

<b> Дополнительно:</b>
• <b>/unset reminder &lt;id&gt; &lt;минуты&gt;</b> — удалить напоминание  
  <i>Пример:</i> <code>/unset</code> <code>reminder</code> <code>123 60</code>

 <code>/events</code> — список событий  
 <code>/showevent &lt;id&gt;</code> — подробности события
`)

	return c.Reply(helpText, &telebot.SendOptions{
		ParseMode: telebot.ModeHTML,
	})
}
