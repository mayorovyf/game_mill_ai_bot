// internal/utils/mode_utils/mode_order.go
package mode_utils

import "game_mill_ai_bot/internal/config"

// задаем уровни типам запуска
func ModeOrder(mode config.Mode) int {
	switch mode {
	case config.DevMode:
		return 1
	case config.TestMode:
		return 2
	case config.ProdMode:
		return 3
	default:
		return 99
	}
}
