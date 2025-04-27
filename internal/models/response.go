package models

type ResponseLevel string

const (
	LevelInfo  ResponseLevel = "info"
	LevelWarn  ResponseLevel = "warn"
	LevelError ResponseLevel = "error"
)

type Response struct {
	Code          int
	Level         ResponseLevel
	Description   string
	VisibleToUser bool
}
