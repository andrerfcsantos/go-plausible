package plausible

import "fmt"

// TimeInterval represents an interval of time by which each entry in the query results must be separated by.
type TimeInterval string

const (
	// DateInterval represents a time interval for a particular day.
	// With this time interval, each data point in the result will refer to a particular day.
	DateInterval = TimeInterval("date")
	// MonthInterval represents a time interval of a particular month.
	// With this time interval, each data point in the result will refer to a particular month.
	MonthInterval = TimeInterval("month")
)

func (t *TimeInterval) toQueryArgs() QueryArgs {
	return QueryArgs{
		{Name: "interval", Value: string(*t)},
	}
}

// IsEmpty tells whether the time interval has information.
func (t *TimeInterval) IsEmpty() bool {
	return string(*t) == ""
}

// TimePeriod represents a period of time for which to get the results.
//
// TimePeriod provides low-level access to the API.
// For most users, it is recommended to use one of the methods in this
// package that returns a time period to build a time period,
// instead of using this struct directly.
type TimePeriod struct {
	// Period is a string representing a period of time, e.g "6mo", "12mo", "7d", "30d", "custom", "month", "day" or "custom".
	// This field is mandatory.
	Period string
	// Date is a string representing a date to which the time period refers to, in the format of "yyyy-mm-dd"
	// This field is optional.
	Date string
}

// Last6Months returns a time period referring to the last 6 months.
// To change the date from which the "last 6 months" refer to,
// chain the return of this function with OfDate() or FromDate() to add
// date information to the time period.
func Last6Months() TimePeriod {
	return TimePeriod{Period: "6mo"}
}

// Last12Months returns a time period referring to the last 12 months.
// To change the date from which the "last 6 months" refer to,
// chain the return of this function with OfDate or FromDate to add
// date information to the time period.
func Last12Months() TimePeriod {
	return TimePeriod{Period: "12mo"}
}

// Last7Days returns a time period referring to the last 7 days.
// To change the date from which the "last 7 days" refer to,
// chain the return of this function with OfDate or FromDate to add
// date information to the time period.
func Last7Days() TimePeriod {
	return TimePeriod{Period: "7d"}
}

// Last30Days returns a time period referring to the last 30 days.
// To change the date from which the "last 30 days" refer to,
// chain the return of this function with OfDate or FromDate to add
// date information to the time period.
func Last30Days() TimePeriod {
	return TimePeriod{Period: "30d"}
}

// MonthPeriod returns a time period referring to a month.
// If no additional date information is given, this defaults to mean "the current month".
// To change the month to which this date refers to,
// chain the return of this function with OfDate or FromDate to add
// date information to the time period.
func MonthPeriod() TimePeriod {
	return TimePeriod{Period: "month"}
}

// DayPeriod returns a time period referring to a day.
// If no additional date information is given, this defaults to mean "today".
// To change the day to which this date refers to,
// chain the return of this function with OfDate or FromDate to add
// date information to the time period.
func DayPeriod() TimePeriod {
	return TimePeriod{Period: "day"}
}

// CustomPeriod allows to build a custom time period/range between the two given dates.
func CustomPeriod(fromDate Date, toDate Date) TimePeriod {
	return TimePeriod{
		Period: "custom",
		Date:   fmt.Sprintf("%s,%s", fromDate.toPlausibleFormat(), toDate.toPlausibleFormat()),
	}
}

// IsEmpty tells whether the time period is empty
func (tp TimePeriod) IsEmpty() bool {
	return tp.Period == ""
}

// OfDate adds date information to a time period. An alias of FromDate.
func (tp TimePeriod) OfDate(date Date) TimePeriod {
	return tp.FromDate(date)
}

// FromDate adds date information to a time period.
func (tp TimePeriod) FromDate(date Date) TimePeriod {
	tp.Date = date.toPlausibleFormat()
	return tp
}

func (tp *TimePeriod) toQueryArgs() QueryArgs {
	qargs := QueryArgs{
		QueryArg{Name: "period", Value: tp.Period},
	}

	if tp.Date != "" {
		qargs.Add(QueryArg{Name: "date", Value: tp.Date})
	}

	return qargs
}
