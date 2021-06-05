package plausible_test

import (
	"fmt"

	"github.com/andrerfcsantos/go-plausible/plausible"
)

func ExampleSite_CurrentVisitors() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// Get the current visitors
	visitors, err := mysite.CurrentVisitors()
	if err != nil {
		// handle error
	}

	fmt.Printf("Site %s has %d current visitors!\n", mysite.ID(), visitors)
}

func ExampleSite_Aggregate() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// Get all visitors for today
	todaysVisitorsQuery := plausible.AggregateQuery{
		Period: plausible.DayPeriod(),
		Metrics: plausible.Metrics{
			plausible.Visitors,
		},
	}

	result, err := mysite.Aggregate(todaysVisitorsQuery)
	if err != nil {
		// handle error
	}
	fmt.Printf("Total visitors of %s today: %d\n", mysite.ID(), result.Visitors)
}

func ExampleSite_Timeseries() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// For each day of the last 7 days of the 1st of February 2021,
	// get the number of visitors and page views.
	tsQuery := plausible.TimeseriesQuery{
		Period: plausible.Last7Days().FromDate(plausible.Date{Day: 1, Month: 2, Year: 2021}),
		Metrics: plausible.Metrics{
			plausible.Visitors,
			plausible.PageViews,
		},
		Interval: plausible.DateInterval,
	}

	queryResults, err := mysite.Timeseries(tsQuery)
	if err != nil {
		// handle error
	}

	// Iterate over the data points
	for _, stat := range queryResults {
		fmt.Printf("Date: %s | Visitors: %d | Pageviews: %d\n", stat.Date, stat.Visitors, stat.Pageviews)
	}
}

func ExampleSite_Breakdown() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// For each page, return the number of visitors and page views in the last 7 days
	pageBreakdownQuery := plausible.BreakdownQuery{
		Property: plausible.EventPage,
		Period:   plausible.Last7Days(),
		Metrics: plausible.Metrics{
			plausible.Visitors,
			plausible.PageViews,
		},
	}

	pageBreakdown, err := mysite.Breakdown(pageBreakdownQuery)
	if err != nil {
		// handle error
	}

	for _, stat := range pageBreakdown {
		fmt.Printf("Page: %s | Visitors: %d | Pageviews: %d\n", stat.Page, stat.Visitors, stat.Pageviews)
	}
}

func ExampleSite_SharedLink() {
	// Create a client with an API token
	// Warning: This token must have permissions to the site provisioning API
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	sl := plausible.SharedLinkRequest{
		Name: "Friends Link",
	}
	slResult, err := mysite.SharedLink(sl)

	if err != nil {
		// handle error
	}

	fmt.Printf("Name: %s | URL: %s\n", slResult.Name, slResult.URL)
}

func ExampleClient_CreateNewSite() {
	// Create a client with an API token
	// Warning: This token must have permissions to the site provisioning API
	client := plausible.NewClient("<your_api_token>")

	newSiteRequest := plausible.CreateSiteRequest{
		Domain:   "mynewsite.com",
		Timezone: "Europe/Lisbon",
	}

	// Note that we call CreateNewSite directly on the client,
	// and not on a site like the majority of requests
	siteResult, err := client.CreateNewSite(newSiteRequest)

	if err != nil {
		// handle error
	}
	fmt.Printf("Domain: %s | Timezone: %s\n", siteResult.Domain, siteResult.Timezone)
}

func ExampleLast6Months() {
	// Get the period for the last 6 months from today
	plausible.Last6Months()
}

func ExampleTimePeriod_FromDate() {
	// Get the period for the first 15 days of 2021
	plausible.CustomPeriod(plausible.Date{Day: 1, Month: 1, Year: 2021}, plausible.Date{Day: 15, Month: 1, Year: 2021})
}

func ExampleCustomPeriod() {
	// Get the period for the last 12 months from the 1st of January 2021
	plausible.Last12Months().FromDate(plausible.Date{Day: 1, Month: 1, Year: 2021})
}

func ExampleFilter_example1() {
	// Filter by windows users
	plausible.NewFilter().ByVisitOs("Windows")
}

func ExampleFilter_example2() {
	// Filter by windows users using firefox
	plausible.NewFilter().ByVisitOs("Windows").ByVisitBrowser("Firefox")
}

func ExampleFilter_example3() {
	// Filter by windows and ubuntu users using firefox
	plausible.NewFilter().ByVisitOs("Windows|Ubuntu").ByVisitBrowser("Firefox")
}

func ExampleFilter_example4() {
	// The following are 3 ways to build the same filter

	// Using the builder pattern to chain properties
	_ = plausible.NewFilter().ByVisitOs("Windows").ByVisitBrowser("Firefox")

	// Using NewFilter
	_ = plausible.NewFilter(
		plausible.Property{Name: plausible.VisitOs, Value: "Windows"},
		plausible.Property{Name: plausible.VisitBrowser, Value: "Firefox"},
	)

	// Instantiating a Filter struct directly
	_ = plausible.Filter{
		Properties: plausible.Properties{
			{Name: plausible.VisitOs, Value: "Windows"},
			{Name: plausible.VisitBrowser, Value: "Firefox"},
		}}
}

func ExampleMetrics() {
	// Visitors and page views metrics
	_ = plausible.Metrics{
		plausible.Visitors,
		plausible.PageViews,
	}
}

func ExampleProperties() {
	// A list of properties with an OS and Browser
	_ = plausible.Properties{
		{Name: plausible.VisitOs, Value: "Windows"},
		{Name: plausible.VisitBrowser, Value: "Firefox"},
	}
}
