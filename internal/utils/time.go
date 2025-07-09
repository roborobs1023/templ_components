package utils

import "time"

var now = time.Now()

// LastDayOfCurrentMonth returns the last date of the current month in a time.Time format.
func LastDayOfCurrentMonth() time.Time {
	// get the first of next month
	firstOfNextMonth := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, now.Location())

	//return the last day of the month calculated as the first of next month -1 day.
	return firstOfNextMonth.AddDate(0, 0, -1)
}

func LastDayOfCurrentYear() time.Time {

	return time.Date(now.Year(), time.December, 31, 0, 0, 0, 0, now.Location())
}

// FirstDayOfCurrentMonth returns the first date of the current month in a time.Time format.
func FirstDayOfCurrentMonth() time.Time {

	return time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
}

func FirstDayOfCurrentYear() time.Time {

	return time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, now.Location())
}

func FirstDayOfYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
}

func FirstOfMonth(month time.Month, year int) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
}

func LastDayOfMonth(month time.Month, year int) time.Time {
	firstOfFollowingMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, now.Location())

	return firstOfFollowingMonth.AddDate(0, 0, -1)
}
