package models

import "game_mill_ai_bot/internal/config"

type User struct {
	ID           string             `json:"id" bson:"id"`
	Username     string             `json:"username" bson:"username"`
	Cloudlets    int                `json:"cloudlets" bson:"cloudlets"`
	Adminlvl     int                `json:"Adminlvl" bson:"Adminlvl"`
	Access       config.AccessLevel `json:"access" bson:"access"`
	CurrentModel string             `json:"CurrentModel" bson:"CurrentModel"`
}
