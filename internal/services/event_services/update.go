package event_services

import (
	"errors"
	"game_mill_ai_bot/internal/db/repository/r_event"
	"time"
)

func UpdateTitle(userID int64, localID int, title string) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.Title = title
	return r_event.UpdateEvent(event)
}

func UpdateDescription(userID int64, localID int, descr string) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.Description = descr
	return r_event.UpdateEvent(event)
}

func UpdateTime(userID int64, localID int, t time.Time) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.StartTime = t
	return r_event.UpdateEvent(event)
}

func AddReminder(userID int64, localID int, mins int) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	for _, m := range event.ReminderMins {
		if m == mins {
			return errors.New("такое напоминание уже есть")
		}
	}
	event.ReminderMins = append(event.ReminderMins, mins)
	return r_event.UpdateEvent(event)
}

func RemoveReminder(userID int64, localID int, mins int) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	newArr := make([]int, 0, len(event.ReminderMins))
	for _, m := range event.ReminderMins {
		if m != mins {
			newArr = append(newArr, m)
		}
	}
	event.ReminderMins = newArr
	return r_event.UpdateEvent(event)
}

func SetTopic(userID int64, localID int, topicID int64) error {
	event, err := r_event.FindEventByLocalID(userID, localID)
	if err != nil {
		return err
	}
	event.TopicID = &topicID
	return r_event.UpdateEvent(event)
}
