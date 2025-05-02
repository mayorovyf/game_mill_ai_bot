package utils

import "game_mill_ai_bot/internal/models"

func isValidLanguage(lang models.Language) bool {
	return lang == models.LangRU || lang == models.LangEN
}
