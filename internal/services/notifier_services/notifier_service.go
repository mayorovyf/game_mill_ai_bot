package notifier_services

import (
	"fmt"
	"gopkg.in/telebot.v3"
)

// NotifierService — сервис отправки напоминаний
type NotifierService struct {
	bot *telebot.Bot
}

// NewNotifierService — конструктор
func NewNotifierService(bot *telebot.Bot) *NotifierService {
	return &NotifierService{bot: bot}
}

// SendEventReminder отправляет напоминание в группу (и топик, если есть)
func (s *NotifierService) SendEventReminder(chatID int64, topicID *int64, text string) error {
	opts := &telebot.SendOptions{}
	if topicID != nil {
		opts.ThreadID = int(*topicID)
	}
	_, err := s.bot.Send(telebot.ChatID(chatID), text, opts)
	if err != nil {
		return fmt.Errorf("ошибка отправки напоминания в чат %d (топик %v): %w", chatID, topicID, err)
	}
	return nil
}
