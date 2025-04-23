package telegram

import (
	"game_mill_ai_bot/internal/telegram/handlers/admin_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/ai_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/event_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/main_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/team_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/user_handlers"
)

func LoadTgRoutes() {

	// main
	bot.Handle("/start", main_handlers.StartHandler)

	// ии
	bot.Handle("/ai", ai_handlers.AiHendler)

	// пользователь
	bot.Handle("/profile", user_handlers.ProfileHandler)

	// админ
	bot.Handle("/ch", admin_handlers.ChangeCloudletsHandler)

	// команды
	bot.Handle("/create_team", team_handlers.CreateTeamHandler)
	bot.Handle("/team_info", team_handlers.TeamInfoHandler)
	bot.Handle("/set_team_name", team_handlers.SetTeamNameHandler)
	bot.Handle("/add_to_team", team_handlers.AddToTeamHandler)

	// события
	bot.Handle("/create_event", event_handlers.CreateEventHandler)
	bot.Handle("/set_event", event_handlers.SetEventFieldHandler)
	bot.Handle("/event_info", event_handlers.EventInfoHandler)
}
