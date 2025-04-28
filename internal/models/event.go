// internal/models/event.go
package models

import "time"

type EventStatus string

const (
	StatusDraft    EventStatus = "draft"
	StatusReady    EventStatus = "ready"
	StatusArchived EventStatus = "archived"
)

// Событие
type Event struct {
	GlobalID      int64           `bson:"global_id"`
	LocalID       int             `bson:"local_id"`
	AuthorID      int64           `bson:"author_id"`
	Title         string          `bson:"title"`
	Description   string          `bson:"description"`
	StartTime     time.Time       `bson:"start_time"`
	ReminderMins  []int           `bson:"reminder_mins"`
	Started       bool            `bson:"started"`
	Status        EventStatus     `bson:"status"`
	ChatID        int64           `bson:"chat_id"`
	TopicID       *int64          `bson:"topic_id,omitempty"`
	Subscribers   []int64         `bson:"subscribers"`
	RemindersSent map[string]bool `bson:"reminders_sent"`
}
