package event_services

import "game_mill_ai_bot/internal/db/repository/r_event"

func DeleteEvent(userID int64, localID int) error {
	return r_event.DeleteEvent(userID, localID)
}
