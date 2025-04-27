// internal/services/response_services/format_message_service.go
package response_services

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/models"
	"game_mill_ai_bot/internal/utils/mode_utils"
	"strings"
)

// формируем сообщение в зависимости от режима
func FormatMessage(resp models.Response) string {
	// выдаём пустую строку в случае если сообщение не видно пользователю или не подходящий режим запуска
	if !resp.VisibleToUser || mode_utils.ModeOrder(config.CurrentMode) < mode_utils.ModeOrder(resp.MinVisibleMode) {
		return ""
	}

	// создаём основу
	var sb strings.Builder
	sb.WriteString(resp.Description)

	switch config.CurrentMode {
	case config.DevMode:
		// в dev моде прописываем подробные детали и техническую информацию
		if resp.UserDetails != "" {
			sb.WriteString("\n\n")
			sb.WriteString(resp.UserDetails)
		}
		if resp.InternalDetails != "" {
			sb.WriteString("\n\n[DEBUG]: ")
			sb.WriteString(resp.InternalDetails)
		}
	case config.TestMode:
		// в test моде прописываем только подробные детали
		if resp.UserDetails != "" {
			sb.WriteString("\n\n")
			sb.WriteString(resp.UserDetails)
		}
	case config.ProdMode:
		// только Description
		// всё нужное уже есть в основе
	}

	// возвращаем строку
	return sb.String()
}
