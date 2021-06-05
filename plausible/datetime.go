package plausible

import "fmt"

// DateTime contains basic information about a date and a time.
// It represents the information about date and time returned by some API calls.
//
// This basic information about the date includes the day of month, month, year.
// The information about the time includes the hour, minute and second.
// It does not contain any information about timezones.
type DateTime struct {
	Date
	Time
}

// String converts the DateTime into a human-readable string.
func (dt *DateTime) String() string {
	return dt.toPlausibleFormat()
}

// toPlausibleFormat converts the DateTime to a string format that the plausible API uses.
func (dt *DateTime) toPlausibleFormat() string {
	return fmt.Sprintf("%s %s", dt.Date.toPlausibleFormat(), dt.Time.toPlausibleFormat())
}

// Time represents the basic information about a time of day.
// It represents the information about time returned by some API calls.
type Time struct {
	// Hour of a given time of day
	Hour int
	// Minute of a given time of day
	Minute int
	// Second of a given time of day
	Second int
}

// String converts the Time into a human-readable string.
func (t *Time) String() string {
	return t.toPlausibleFormat()
}

// toPlausibleFormat converts the Time to a string format that the plausible API uses.
func (t *Time) toPlausibleFormat() string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

// Date represents the basic information about a date.
// It represents the information about date returned by some API calls.
type Date struct {
	// Day of month of the date (1-31)
	Day int
	// Month of the date (1-12)
	Month int
	// Year of the date (1-12)
	Year int
}

// String converts the Date into a human-readable string.
func (d *Date) String() string {
	return d.toPlausibleFormat()
}

// toPlausibleFormat converts the Date to a string format that the plausible API uses.
func (d *Date) toPlausibleFormat() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}
