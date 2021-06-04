package plausible

import "testing"

func TestUnitTime(t *testing.T) {

	tests := []struct {
		name                    string
		time                    Time
		expectedString          string
		expectedPlausibleFormat string
	}{
		{
			name: "valid time with single digits for hours, minutes and seconds",
			time: Time{
				Hour:   1,
				Minute: 2,
				Second: 3,
			},
			expectedString:          "01:02:03",
			expectedPlausibleFormat: "01:02:03",
		},
		{
			name: "valid time with double digits for hours, minutes and seconds",
			time: Time{
				Hour:   10,
				Minute: 20,
				Second: 30,
			},
			expectedString:          "10:20:30",
			expectedPlausibleFormat: "10:20:30",
		},
	}

	for _, test := range tests {
		if test.time.String() != test.expectedString {
			t.Fatalf("test '%s' failed: expected conversion to string to be %s, got %s",
				test.name, test.expectedString, test.time.String(),
			)
		}

		if test.time.toPlausibleFormat() != test.expectedPlausibleFormat {
			t.Fatalf("test '%s' failed: expected conversion to plausible format to be %s, got %s",
				test.name, test.expectedPlausibleFormat, test.time.toPlausibleFormat(),
			)
		}
	}

}

func TestUnitDate(t *testing.T) {

	tests := []struct {
		name                    string
		date                    Date
		expectedString          string
		expectedPlausibleFormat string
	}{
		{
			name: "valid date with single digits for day and month",
			date: Date{
				Day:   1,
				Month: 2,
				Year:  2021,
			},
			expectedString:          "2021-02-01",
			expectedPlausibleFormat: "2021-02-01",
		},
		{
			name: "valid date with double digits for day and month",
			date: Date{
				Day:   10,
				Month: 12,
				Year:  2021,
			},
			expectedString:          "2021-12-10",
			expectedPlausibleFormat: "2021-12-10",
		},
	}

	for _, test := range tests {
		if test.date.String() != test.expectedString {
			t.Fatalf("test '%s' failed: expected conversion to string to be %s, got %s",
				test.name, test.expectedString, test.date.String(),
			)
		}

		if test.date.toPlausibleFormat() != test.expectedPlausibleFormat {
			t.Fatalf("test '%s' failed: expected conversion to plausible format to be %s, got %s",
				test.name, test.expectedPlausibleFormat, test.date.toPlausibleFormat(),
			)
		}
	}

}

func TestUnitDateTime(t *testing.T) {

	tests := []struct {
		name                    string
		datetime                DateTime
		expectedString          string
		expectedPlausibleFormat string
	}{
		{
			name: "valid datetime time with single digits for day, month, hours, minutes and seconds",
			datetime: DateTime{
				Date: Date{
					Day:   1,
					Month: 2,
					Year:  2021,
				},
				Time: Time{
					Hour:   1,
					Minute: 2,
					Second: 3,
				},
			},
			expectedString:          "2021-02-01 01:02:03",
			expectedPlausibleFormat: "2021-02-01 01:02:03",
		},
		{
			name: "valid datetime time with double digits for day, month, hours, minutes and seconds",
			datetime: DateTime{
				Date: Date{
					Day:   10,
					Month: 12,
					Year:  2021,
				},
				Time: Time{
					Hour:   10,
					Minute: 20,
					Second: 30,
				},
			},
			expectedString:          "2021-12-10 10:20:30",
			expectedPlausibleFormat: "2021-12-10 10:20:30",
		},
	}

	for _, test := range tests {
		if test.datetime.String() != test.expectedString {
			t.Fatalf("test '%s' failed: expected conversion to string to be %s, got %s",
				test.name, test.expectedString, test.datetime.String(),
			)
		}

		if test.datetime.toPlausibleFormat() != test.expectedPlausibleFormat {
			t.Fatalf("test '%s' failed: expected conversion to plausible format to be %s, got %s",
				test.name, test.expectedPlausibleFormat, test.datetime.toPlausibleFormat(),
			)
		}
	}

}
