package utils

import (
	"fmt"
	"strconv"
)

func parseTimezone(arg string) (int, error) {
	tz, err := strconv.Atoi(arg)
	if err != nil || tz < -12 || tz > 14 {
		return 0, fmt.Errorf("Неверный формат. Используйте целое число от -12 до +14.")
	}
	return tz, nil
}
