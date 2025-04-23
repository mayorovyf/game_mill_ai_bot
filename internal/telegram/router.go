package telegram

import (
	"game_mill_ai_bot/internal/telegram/handlers"
	"game_mill_ai_bot/internal/telegram/handlers/event_handlers"
)

func LoadTgRoutes() {
	bot.Handle("/start", handlers.StartHandler)
	bot.Handle("/ai", handlers.AiHendler)
	bot.Handle("/profile", handlers.ProfileHandler)
	bot.Handle("/ch", handlers.ChangeCloudletsHandler)
	bot.Handle("/create_team", handlers.CreateTeamHandler)
	bot.Handle("/team_info", handlers.TeamInfoHandler)
	bot.Handle("/set_team_name", handlers.SetTeamNameHandler)
	bot.Handle("/add_to_team", handlers.AddToTeamHandler)
	bot.Handle("/create_event", event_handlers.CreateEventHandler)
	bot.Handle("/set_event", event_handlers.SetEventFieldHandler)
}
