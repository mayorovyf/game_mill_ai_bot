package access

import "game_mill_ai_bot/internal/config"

var accessOrder = map[config.AccessLevel]int{
	config.AccessUser:       0,
	config.AccessModerator:  1,
	config.AccessAdmin:      2,
	config.AccessSuperAdmin: 3,
}

// HasAccess возвращает true, если у пользователя доступ >= требуемого
func HasAccess(userLevel, required config.AccessLevel) bool {
	return accessOrder[userLevel] >= accessOrder[required]
}
