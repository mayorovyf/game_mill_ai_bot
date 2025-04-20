package models

type User struct {
	ID           string `json:"id" bson:"id"`
	Username     string `json:"username" bson:"username"`
	Cloudlets    int    `json:"cloudlets" bson:"cloudlets"`
	Adminlvl     int    `json:"Adminlvl" bson:"Adminlvl"`
	CurrentModel string `json:"CurrentModel" bson:"CurrentModel"`
}
