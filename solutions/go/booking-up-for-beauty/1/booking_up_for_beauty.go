package booking

import (
	"fmt"
	"time"
)

const (
	layout1 = "1/2/2006 15:04:05"
	layout2 = "January 2, 2006 15:04:05"
	layout3 = "Monday, January 2, 2006 15:04:05"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// => 2019-07-25 13:45:00 +0000 UTC

	layouts := []string{layout1, layout2, layout3}
	for _, lay := range layouts {
		if scheduled, err := time.Parse(lay, date); err == nil {
			return scheduled
		}
	}
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// Compare compares the time instant t with u. If t is before u, it returns -1;
	// if t is after u, it returns +1; if they're the same, it returns 0.
	return time.Now().Compare(Schedule(date)) == +1
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	scheduled := Schedule(date)
	// Hour() returns an int in [0, 23] range.
	return scheduled.Hour() >= 12 && scheduled.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	scheduled := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", scheduled.Weekday().String(), scheduled.Month().String(), scheduled.Day(), scheduled.Year(), scheduled.Hour(), scheduled.Minute())
	// => "You have an appointment on Thursday, July 25, 2019, at 13:45."

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	
	anniversary := fmt.Sprintf("09/15/%d 00:00:00", time.Now().Year())
	return Schedule(anniversary)
}
