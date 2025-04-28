package notifier

import (
	"fmt"
	"strconv"
	"time"

	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/services/notifier_services"
	"gopkg.in/telebot.v3"
)

func StartEventNotifier(bot *telebot.Bot) {
	notifier := notifier_services.NewNotifierService(bot)
	go func() {
		ticker := time.NewTicker(time.Minute)
		defer ticker.Stop()
		for {
			<-ticker.C
			now := time.Now().UTC()
			events, err := r_event.FindUpcomingEvents(120)
			if err != nil {
				continue
			}
			fmt.Printf("Найдено событий для проверки: %d\n", len(events))
			for _, ev := range events {
				for _, mins := range ev.ReminderMins {
					notifyTime := ev.StartTime.Add(-time.Duration(mins) * time.Minute)
					if now.After(notifyTime.Add(-time.Second)) && now.Before(notifyTime.Add(time.Minute)) {
						minsStr := strconv.Itoa(mins)
						if ev.RemindersSent == nil {
							ev.RemindersSent = make(map[string]bool)
						}
						if !ev.RemindersSent[minsStr] {
							msg := fmt.Sprintf(
								"Напоминание: событие \"%s\" начнётся через %d минут!",
								ev.Title, mins,
							)
							err := notifier.SendEventReminder(ev.ChatID, ev.TopicID, msg)
							if err != nil {
								fmt.Println("Ошибка отправки напоминания:", err)
							}
							ev.RemindersSent[minsStr] = true
							_ = r_event.UpdateEventRemindersSent(ev.GlobalID, ev.RemindersSent)
						}
					}
				}
			}
		}
	}()
}
