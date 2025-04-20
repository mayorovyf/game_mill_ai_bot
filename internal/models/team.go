package models

type Team struct {
	Name    string   `json:"name" bson:"name"`
	Id      string   `json:"id" bson:"id"`
	ChatId  string   `json:"chatId" bson:"chatId"`
	Members []string `json:"members" bson:"members"`
	Admins  []string `json:"admins" bson:"admins"`
	Lvl     int      `json:"lvl" bson:"lvl"`
}
