package utils

import (
	"game_mill_ai_bot/internal/config"
	"game_mill_ai_bot/internal/models"
)

// Error возвращает стандартный error-ответ
func Error(desc, userDesc string, err error) models.Response {
	return models.Response{
		Level:           models.LevelError,
		Description:     desc,
		UserDetails:     userDesc,
		InternalDetails: err.Error(),
		VisibleToUser:   true,
		MinVisibleMode:  config.TestMode,
	}
}

// Info возвращает стандартный информационный ответ
func Info(desc, userDesc string) models.Response {
	return models.Response{
		Level:          models.LevelInfo,
		Description:    desc,
		UserDetails:    userDesc,
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}

// Warn возвращает предупреждающий ответ
func Warn(desc, userDesc string) models.Response {
	return models.Response{
		Level:          models.LevelWarn,
		Description:    desc,
		UserDetails:    userDesc,
		VisibleToUser:  true,
		MinVisibleMode: config.ProdMode,
	}
}
