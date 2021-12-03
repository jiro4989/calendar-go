/*
Calendar is tiny library to create calendar.

*/
package calendar

import (
	"time"
)

// CalendarDays return days like `cal` calendar.
//
// The days like `cal` calendar are:
//
//      December 2021
//   Su Mo Tu We Th Fr Sa
//   28 29 30  1  2  3  4
//    5  6  7  8  9 10 11
//   12 13 14 15 16 17 18
//   19 20 21 22 23 24 25
//   26 27 28 29 30 31  1
func CalendarDays(y int, m time.Month) []time.Time {
	b := beginDay(y, m)
	e := endDay(y, m)

	var days []time.Time
	weekDiff := b.Weekday() - time.Sunday
	for i := int(weekDiff); 0 < i; i-- {
		d := b.AddDate(0, 0, -i)
		days = append(days, d)
	}

	diff := e.Day() - b.Day()
	for i := 0; i < diff; i++ {
		d := b.AddDate(0, 0, i)
		days = append(days, d)
	}

	weekDiff = time.Saturday - e.Weekday()
	for i := 0; i <= int(weekDiff); i++ {
		d := e.AddDate(0, 0, i)
		days = append(days, d)
	}

	return days
}

func beginDay(y int, m time.Month) time.Time {
	return time.Date(y, m, 1, 0, 0, 0, 0, time.Local)
}

func endDay(y int, m time.Month) time.Time {
	n := nextMonth(y, m)
	n = n.AddDate(0, 0, -1)
	return n
}

func nextMonth(y int, m time.Month) time.Time {
	if m == time.December {
		y++
		m = time.January
	} else {
		m++
	}
	return beginDay(y, m)
}
