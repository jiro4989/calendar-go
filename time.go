package calendar

import "time"

// IsSameDay returns true when a and b are same day.
//
// This function ignores hour, minute, seconds.
func IsSameDay(a, b time.Time) bool {
	return a.Year() == b.Year() && a.Month() == b.Month() && a.Day() == b.Day()
}

// ContainsSameDay returns bool and found day.
//
// This function ignores hour, minute, seconds.
func ContainsSameDay(day time.Time, days []time.Time) bool {
	for _, d := range days {
		if IsSameDay(day, d) {
			return true
		}
	}
	return false
}
