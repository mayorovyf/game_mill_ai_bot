package utils

import (
	"errors"
	"strings"
	"time"
)

func ParseTime(input string, offsetHours int) (time.Time, error) {
	input = strings.TrimSpace(input)

	// текущее время в заданной временной зоне
	location := time.FixedZone("custom", offsetHours*3600)
	now := time.Now().In(location)
	year := now.Year()

	formats := []string{
		"15",               // hh
		"15:04",            // hh:mm
		"02 15:04",         // dd hh:mm
		"02.01 15:04",      // dd.mm hh:mm
		"02.01.06 15:04",   // dd.mm.yy hh:mm
		"02.01.2006 15:04", // dd.mm.yyyy hh:mm
		"15:04 02.01.2006", // hh:mm dd.mm.yyyy
		"15:04 02",         // hh:mm dd
		"15:04 02.01",      // hh:mm dd.mm
	}

	for _, layout := range formats {
		t, err := time.ParseInLocation(layout, input, location)
		if err == nil {
			switch layout {
			case "15":
				return time.Date(year, now.Month(), now.Day(), t.Hour(), 0, 0, 0, location).UTC(), nil

			case "02 15:04":
				dt := time.Date(year, now.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, location)
				if !dt.After(now) {
					dt = dt.AddDate(0, 0, 7)
				}
				return dt.UTC(), nil

			case "15:04 02":
				dt := time.Date(year, now.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, location)
				if !dt.After(now) {
					dt = dt.AddDate(0, 0, 1)
				}
				return dt.UTC(), nil

			case "15:04 02.01":
				dt := time.Date(year, t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, location)
				if !dt.After(now) {
					dt = dt.AddDate(1, 0, 0)
				}
				return dt.UTC(), nil

			default:
				return t.UTC(), nil
			}
		}
	}

	return time.Time{}, errors.New("Неверный формат времени. Примеры: 18, 18:00, 05.06 18:00")
}
