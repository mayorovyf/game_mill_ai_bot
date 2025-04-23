package event_handlers

import (
	"game_mill_ai_bot/internal/db/repository/r_event"
	"game_mill_ai_bot/internal/models"
	"gopkg.in/telebot.v3"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func CreateEventHandler(c telebot.Context) error {
	// Генерируем 10-значный числовой ID
	rand.Seed(time.Now().UnixNano())
	idNum := rand.Int63n(9_000_000_000) + 1_000_000_000
	eventID := strconv.FormatInt(idNum, 10)

	event := &models.Event{
		ID:           eventID,
		ChatID:       strconv.FormatInt(c.Chat().ID, 10),
		Title:        "Title",           // title
		Description1: "Description1",    // description1
		Description2: "Description2",    // description2
		TypeID:       "default_type_id", // type_id
		Type:         "default_type",    // type
		Date:         time.Now(),        // date
		Subscribers:  []string{},
	}
	if err := r_event.CreateEvent(event); err != nil {
		return c.Reply("Ошибка при создании события: " + err.Error())
	}

	// Кнопка подписки
	btn := telebot.InlineButton{
		Unique: "subscribe_event_" + eventID,
		Text:   "✅ Подписаться",
	}
	c.Bot().Handle(&btn, func(c telebot.Context) error {
		return SubscribeToEvent(c, eventID)
	})

	// Выводим дефолты и подсказки по тэгам (без #)
	var sb strings.Builder
	sb.WriteString("✅ Событие создано! (ID: " + eventID + ")\n\n")
	sb.WriteString("Title: " + event.Title + " (`title`)\n")
	sb.WriteString("Description1: " + event.Description1 + " (`description1`)\n")
	sb.WriteString("Description2: " + event.Description2 + " (`description2`)\n")
	sb.WriteString("TypeID: " + event.TypeID + " (`type_id`)\n")
	sb.WriteString("Type: " + event.Type + " (`type`)\n")
	sb.WriteString("Date: " + event.Date.Format("2006-01-02 15:04") + " (`date`)\n\n")
	sb.WriteString("Чтобы изменить поле, используй:\n")
	sb.WriteString("/set_event " + eventID + " tag новое_значение\n")
	sb.WriteString("Пример: /set_event " + eventID + " title НовоеНазвание\n")

	return c.Reply(sb.String(), &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{{btn}},
	})
}
