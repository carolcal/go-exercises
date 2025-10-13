package booking

import (
    "fmt"
    "time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	formats := []string{
        "1/2/2006 15:04:05",
        "January 2, 2006 15:04:05",
        "Monday, January 2, 2006 15:04:05",
    }
    for _, format := range formats {
		t, err := time.Parse(format, date)
        if err == nil {
            return t
        }
    }
    return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    h := Schedule(date).Hour()
    if h >= 12 && h <= 18 {
        return true
    }
    return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t := Schedule(date)
    return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    currentYear := time.Now().Year()
    anniversary := fmt.Sprintf("09/15/%d 00:00:00", currentYear)
	return Schedule(anniversary)
}
