package chat_handlers

import (
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/services/chat_services"
	"gopkg.in/telebot.v3"
	"strings"
)

// /setlanguage ru
func SetLanguageHandler(c telebot.Context) error {
	arg := strings.TrimSpace(c.Message().Payload)
	lang := models.Language(arg)

	if lang != models.LangRU && lang != models.LangEN {
		return c.Reply("Допустимые языки: ru, en")
	}

	if err := chat_services.SetChatLanguage(c.Chat().ID, lang); err != nil {
		return c.Reply("Ошибка установки языка.")
	}

	return c.Reply("Язык чата обновлён: " + string(lang))
}
