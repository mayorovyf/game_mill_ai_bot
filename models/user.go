package models

type User struct {
	ID        string `json:"id" bson:"id"`
	Username  string `json:"username" bson:"username"`
	Cloudlets string `json:"cloudlets" bson:"cloudlets"`
}
