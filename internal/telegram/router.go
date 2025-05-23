package telegram

import (
	"game_mill_ai_bot/internal/telegram/handlers/admin_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/ai_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/event_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/event_handlers/help"
	"game_mill_ai_bot/internal/telegram/handlers/group_admin_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/main_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/team_handlers"
	"game_mill_ai_bot/internal/telegram/handlers/user_handlers"
)

func LoadTgRoutes() {
	//  Общие
	bot.Handle("/start", main_handlers.StartHandler)
	bot.Handle("/profile", user_handlers.ProfileHandler)
	bot.Handle("/id", main_handlers.IDHandler)

	//  Искусственный интеллект
	bot.Handle("/ai", ai_handlers.AiHendler)

	//  Администрирование
	bot.Handle("/ch", admin_handlers.ChangeCloudletsHandler)

	//  Команды
	bot.Handle("/create_team", team_handlers.CreateTeamHandler)
	bot.Handle("/team_info", team_handlers.TeamInfoHandler)
	bot.Handle("/set_team_name", team_handlers.SetTeamNameHandler)
	bot.Handle("/add_to_team", team_handlers.AddToTeamHandler)

	//  События
	bot.Handle("/newevent", event_handlers.NewEventHandler)
	bot.Handle("/set", event_handlers.SetHandler)
	bot.Handle("/events", event_handlers.ListEventsHandler)
	bot.Handle("/showevent", event_handlers.ShowEventHandler)
	bot.Handle("/delete", event_handlers.DeleteEventHandler)
	bot.Handle("/subscribe", event_handlers.SubscribeHandler)
	bot.Handle("/unsubscribe", event_handlers.UnsubscribeHandler)
	bot.Handle("/ready", event_handlers.ReadyEventHandler)
	bot.Handle("/archive", event_handlers.ArchiveEventHandler)

	// Администрирование групп
	bot.Handle("/setadmin", group_admin_handlers.SetAdminHandler)
	bot.Handle("/removeadmin", group_admin_handlers.RemoveAdminHandler)
	bot.Handle("/setadmintitle", group_admin_handlers.SetAdminTitleHandler)
	bot.Handle("/setadminrights", group_admin_handlers.SetAdminRightsHandler)

	//  Подсказки
	bot.Handle("/adminhelp", group_admin_handlers.AdminHelpHandler)
	bot.Handle("/sethelp", help.SetHelpHandler)
	bot.Handle("/eventhelp", help.EventHelpHandler)
}
