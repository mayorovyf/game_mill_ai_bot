package notifier

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/db/repository/r_user"
	"gopkg.in/telebot.v3"
	"strconv"
	"time"
)

func StartEventNotifier(bot *telebot.Bot) {
	go func() {
		for {
			now := time.Now()
			events, _ := r_event.GetEventsBefore(now)
			for _, evt := range events {
				users, _ := r_user.GetUsersByIds(evt.Subscribers)

				mentions := ""
				for _, u := range users {
					if u.Username != "" {
						mentions += "@" + u.Username + " "
					}
				}

				msg := "🔔 *" + evt.Title + "* начинается!\n" +
					"_TypeID_: " + evt.TypeID + "\n" +
					"_Type_: " + evt.Type + "\n\n" +
					evt.Description1 + "\n" +
					evt.Description2 + "\n\n" +
					"_Подписчики_: " + mentions

				chatID, _ := strconv.ParseInt(evt.ChatID, 10, 64)
				bot.Send(&telebot.Chat{ID: chatID}, msg, &telebot.SendOptions{
					ParseMode: telebot.ModeMarkdown,
				})
			}
			time.Sleep(1 * time.Minute)
		}
	}()
}
