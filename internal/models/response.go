package models

import "game_mill_ai_bot/internal/config"

type ResponseLevel string

const (
	LevelInfo  ResponseLevel = "info"
	LevelWarn  ResponseLevel = "warn"
	LevelError ResponseLevel = "error"
)

type Response struct {
	Level           ResponseLevel
	Description     string
	UserDetails     string
	InternalDetails string
	VisibleToUser   bool
	MinVisibleMode  config.Mode
}
