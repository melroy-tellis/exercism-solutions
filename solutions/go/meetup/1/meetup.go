package meetup

import "time"

// Define the WeekSchedule type here.
type WeekSchedule string;
const (
    First WeekSchedule = "first"
    Second WeekSchedule = "second"
    Third WeekSchedule = "third"
    Fourth WeekSchedule = "fourth"
    Teenth WeekSchedule = "teenth"
    Last WeekSchedule = "last"
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
    // One-liner to get number of days in the current month
    // See https://brandur.org/fragments/go-days-in-month
    daysInMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, time.UTC).Day()
	var firstDate int
    switch(wSched) {
        case First:
        	firstDate = 1
    	case Second:
        	firstDate = 8
        case Third:
        	firstDate = 15
        case Fourth:
        	fallthrough
        case Last:
            firstDate = 22
    	case Teenth:
            firstDate = 13
    }
    firstWeekday := time.Date(year, month, firstDate, 0, 0, 0, 0, time.UTC).Weekday()
    offset := 0
    if wDay >= firstWeekday {
        offset = int(wDay - firstWeekday)
    } else {
        offset = 7 - int(firstWeekday - wDay)
    }
    date := firstDate + offset
	if wSched == Last && date + 7 <= daysInMonth {
        date += 7
    }
    return date
}
