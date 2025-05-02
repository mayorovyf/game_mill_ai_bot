package models

import (
	"gopkg.in/telebot.v3"
)

type Language string

const (
	LangRU Language = "ru"
	LangEN Language = "en"
)

type Chat struct {
	ID       int64            `bson:"id"`
	Type     telebot.ChatType `bson:"type"`
	Title    string           `bson:"title"`
	Username string           `bson:"username"`
	Timezone int              `bson:"timezone"`
	Language Language         `bson:"language"`
	Admins   []GroupAdmin     `bson:"admins"`
}
