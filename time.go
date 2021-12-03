package calendar

import "time"

func IsSameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

func Contains(day time.Time, days []time.Time) (bool, *time.Time) {
	for _, d := range days {
		if IsSameDay(day, d) {
			return true, &d
		}
	}
	return false, nil
}
