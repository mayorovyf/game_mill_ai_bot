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
	bot.Handle("/newevent", event_handlers.NewEventHandler)
	bot.Handle("/set", event_handlers.SetHandler)
	bot.Handle("/events", event_handlers.ListEventsHandler)
	bot.Handle("/showevent", event_handlers.ShowEventHandler)
	bot.Handle("/delete", event_handlers.DeleteEventHandler)
	bot.Handle("/subscribe", event_handlers.SubscribeHandler)
	bot.Handle("/unsubscribe", event_handlers.UnsubscribeHandler)
	bot.Handle("/sethelp", event_handlers.SetHelpHandler)
	bot.Handle("/ready", event_handlers.ReadyEventHandler)
	bot.Handle("/eventhelp", event_handlers.EventHelpHandler)

}
