package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Парсит строку вида "2025-04-28 23:56 UTC+3" в time.Time (UTC)
func ParseEventTime(input string) (time.Time, error) {
	input = strings.TrimSpace(input)
	parts := strings.Fields(input)
	if len(parts) < 3 {
		return time.Time{}, errors.New("Формат: YYYY-MM-DD HH:MM UTC+N")
	}
	date := parts[0]
	hm := parts[1]
	tz := parts[2]

	// Определим смещение
	offset := 0
	if strings.HasPrefix(tz, "UTC") {
		sign := 1
		zone := tz[3:]
		if len(zone) > 0 {
			switch zone[0] {
			case '+':
				sign = 1
				zone = zone[1:]
			case '-':
				sign = -1
				zone = zone[1:]
			}
			if zone != "" {
				hours, err := strconv.Atoi(zone)
				if err != nil {
					return time.Time{}, errors.New("Неверный формат часового пояса")
				}
				offset = sign * hours
			}
		}
	}

	// Парсим как будто в этом поясе, потом переводим в UTC
	layout := "2006-01-02 15:04"
	local, err := time.Parse(layout, date+" "+hm)
	if err != nil {
		return time.Time{}, errors.New("Формат времени: YYYY-MM-DD HH:MM UTC+N")
	}
	// Переводим в UTC
	utcTime := local.Add(-time.Duration(offset) * time.Hour).UTC()
	return utcTime, nil
}
