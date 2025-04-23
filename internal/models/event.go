package models

import "time"

type Event struct {
	ID           string    `bson:"id"`
	ChatID       string    `bson:"chat_id"`
	Title        string    `bson:"title"`
	Description1 string    `bson:"description_1"`
	Description2 string    `bson:"description_2"`
	TypeID       string    `bson:"type_id"`
	Type         string    `bson:"type"`
	Date         time.Time `bson:"date"`
	Subscribers  []string  `bson:"subscribers"`
}
