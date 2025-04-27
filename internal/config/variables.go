package config

type Mode string

const (
	DevMode  Mode = "dev"
	TestMode Mode = "test"
	ProdMode Mode = "prod"
)

var CurrentMode = ProdMode
