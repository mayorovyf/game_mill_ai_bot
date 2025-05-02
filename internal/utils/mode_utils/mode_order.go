// internal/utils/mode_utils/mode_order.go
package mode_utils

import "game_mill_ai_bot/internal/config"

// задаем уровни типам запуска
func ModeOrder(mode config.Mode) int {
	switch mode {
	case config.DevMode:
		return 99
	case config.TestMode:
		return 98
	case config.ProdMode:
		return 97
	default:
		return 1
	}
}
