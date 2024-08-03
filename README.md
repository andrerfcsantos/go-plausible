# <a name="top"></a> `go-plausible` - Go Wrapper for the Plausible API

[![Go Reference](https://pkg.go.dev/badge/github.com/andrerfcsantos/go-plausible/plausible.svg)](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible) [![Go Report Card](https://goreportcard.com/badge/github.com/andrerfcsantos/go-plausible)](https://goreportcard.com/report/github.com/andrerfcsantos/go-plausible)

Go wrapper/client for the [Plausible](https://plausible.io/) API.

It currently supports the full API of Plausible, which includes:

* [Stats API](https://plausible.io/docs/stats-api)
* [Site Provisioning API](https://plausible.io/docs/sites-api)

## Table of Contents

* [Basic Usage](#basic-usage)

* [Concepts](#concepts)
    * [Time Periods](#time-periods)
    * [Properties](#properties)
    * [Filters](#filters)
    * [Metrics](#metrics)
    * [Time Intervals](#metrics)

* [Queries (Stats API)](#queries)
    * [Current Visitors](#currrent-visitors)
    * [Aggregate Queries](#aggregate-queries)
    * [Time series Queries](#timeseries-queries)
    * [Breakdown Queries](#breakdown-queries)

* [Site Provisioning API](#site-provisioning-api)
    * [List sites](#provisioning-api-get-sites)
    * [Get site](#provisioning-api-get-site)
    * [Get/Create Shared Links](#provisioning-api-shared-links)
    * [Create new sites](#provisioning-api-create-new-sites)

* [Events API](#events-api)

* [Tests](#tests)
    * [Unit Tests](#unit-tests)
    * [Integration Tests](#integration-tests)
    * [Integration Tests with provisioning API](#integration-tests-provisioning)

* [Bugs and Feedback](#bugs-feedback)
* [Contributing](#contributing)
* [License](#license)

## <a name="basic-usage"></a> Basic Usage

```go
import "github.com/andrerfcsantos/go-plausible/plausible"
```

To use this client, you'll need an API token, which you can get from the Plausible Dashboard.

With the API token, create a client and get a handler for one or more sites:

```go
package main

import "github.com/andrerfcsantos/go-plausible/plausible"

func main() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// You can reuse the same client to get handlers for additional sites
	myothersite := client.Site("otherexample.com")
}
```

## <a name="concepts"></a> Concepts

There a few concepts that are useful to know before using this wrapper or the Plausible API.

### <a name="time-periods"></a> Time Periods

When requesting aggregate information to the API, it's only possible to get data for a given period of time. For
instance,
"the last 7 days", "the last month" or "the current day" are examples of time periods.

All time periods are relative to a date. When the date information is missing from a time period, the date is assumed to
be "today". It's also possible to specify a time period between two specific dates.

Time periods are represented in this library by
the [TimePeriod](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#TimePeriod) type.

Unless you want low-level access to the API, you don't need to create a `TimePeriod` directly, and you can just use the
helper functions to build a time period:

```go
// Get the period for the last 6 months from today
p := plausible.Last6Months()
```

To associate a date to a time period, chain the result with `FromDate()` or
`OfDate()`:

```go
// Get the period for the last 12 months from the 1st of January 2021
p := plausible.Last12Months().FromDate(plausible.Date{Day:1, Month: 1, Year: 2021})
```

To make a custom period between 2 dates:

```go
// Get the period for the first 15 days of 2021
p := plausible.CustomPeriod(plausible.Date{Day:1, Month: 1, Year: 2021}, plausible.Date{Day:15, Month: 1, Year: 2021})
```

To know more about time periods, see [Plausible Docs: Time Periods](https://plausible.io/docs/stats-api#time-periods)

### <a name="properties"></a> Properties

Each pageview or custom event has some properties associated with it. These properties can be used when querying the API
to filter the results.

Properties are represented in this library by
the [Property](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#Property)
type. Properties have a name and value. The name of a property is represented by
the [PropertyName](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#PropertyName)
type. Typically, most users won't need to make custom property names and can just use the constant `PropertyName`
values declared at the top-level of the package like `VisitOs` or `VisitBrowser`.

To make a custom `PropertyName`, the function `CustomPropertyName()` can be used:

```go
pName := plausible.CustomPropertyName("myevent")
```

Obtaining custom property names via this method is needed when you have custom events and want to refer to those events
as a property.

To easily make a custom property with a name and a value, you can use the `CustomProperty` function:

```go
p := plausible.CustomProperty("myevent", "myeventvalue")
```

To know more about properties, see [Plausible Docs: Properties](https://plausible.io/docs/stats-api#properties)

### <a name="filters"></a> Filters

Filters allow drilling down and segment the data to which the results refer to. All queries for data accept an optional
filter argument.

In this library, filters are represented by
the [Filter](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#Filter) type.

A filter consists of a simple list of properties by which you want to filter. For instance, to create a filter that
filters all visits by their operating system, you can do:

```go
f := plausible.NewFilter().ByVisitOs("Windows")
```

You can add more properties to the filter by chaining calls:

```go
f := plausible.NewFilter().ByVisitOs("Windows").ByVisitBrowser("Firefox")
```

This will filter the results based on the visits from Windows users that were using Firefox. So, a filter basically
consists of a logic AND of all its properties.

You can also instantiate a filter directly if you want a more low-level access to the API. For instance, this an
alternative way to write the filter above:

```go
f := plausible.Filter{
  Properties: plausible.Properties{
    {Name:  plausible.VisitOs, Value: "Windows"},
    {Name:  plausible.VisitBrowser, Value: "Firefox"},
  }
}
```

For each property, you can provide a set of values, separated by `|` to make the filter match any of the provided
values. For instance, to filter the data by visits of users using firefox in either linux or windows, we can do:

```go
f := plausible.NewFilter().ByVisitOs("Windows|Linux").ByVisitBrowser("Firefox")
```

To know more about properties, see [Plausible Docs: Filtering](https://plausible.io/docs/stats-api#filtering)

### <a name="metrics"></a> Metrics

Metrics are aggregate information about the data. All queries have the option for you to choose the metrics you want to
see included in the results.

There are 6 metrics currently that you can ask the results for: number of visitors, number of page views, visit duration, bounce rate, visits and events.
In this library, these metrics are represented by
the [Metric](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#Metric)
type. There are 6 constants of type `Metric`, each one representing one of the 6 metrics: `Visitors`,
`PageViews`, `BounceRate`, `VisitDuration`, `Events` and `Visits`

For instance, if for a query you only want information about the pageviews and number of visitors, you can pass this to
the query in the metrics parameter:

```go
metrics := plausible.Metrics {
	plausible.Visitors,
	plausible.PageViews,
},
```

For convenience, when you want to get information about all metrics, there's a function `AllMetrics()`
that returns all the 6 metrics. However, please note that not all queries support requests for all metrics. For that
reason, use requests for all metrics with caution. If you try to use a metric in a query that does not support that
metric, you will get an error message saying which property was at fault.

### <a name="time-intervals"></a> Time Intervals

Time intervals are used for [time series queries](#timeseries-queries) to specify the interval of time between 2
consecutive data points.

A time interval is represented by
the [TimeInterval](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible#TimeInterval) type. There are
currently 2 time intervals: date and month. This library also exposes two `TimeInterval`
constants for these value: `DateInterval` and `MonthInterval` respectively.

A `MonthInterval` means a month of difference between data points. For instance, if you ask for time series data over
the last 6 months with a month interval, this means you will get 6 data points back - 1 for each month.

A `DateInterval`, depending on the query, means a day or an hour of difference between each data point. For instance, if
you ask for time series data over the last 30 days with a date interval, you will get 30 data points back - 1 for each
day. However, with a `DateInterval`, when the period of the time series refers to a day, for instance "today", the data
points will actually have 1 hour of interval between them. You can check the `Date` string field of each data point to
know about which date/hour the data refers to.

## <a name="queries"></a> Queries

There are 4 types of queries supported by the API:

* Current Visitors
* Aggregate Queries
* Timeseries Queries
* Breakdown Queries

### <a name="current-visitors"></a> Current Visitors

This is the most straight forward query - for a given site return the number of current visitors:

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	visitors, err := mysite.CurrentVisitors()
	if err != nil {
		// handle error
	}

	fmt.Printf("Site %s has %d current visitors!\n", mysite.ID(), visitors)
}
```

### <a name="aggregate-queries"></a> Aggregate Queries

An aggregate query reports data for metrics aggregated over a period of time.

A query like "the total number of visitors today" fall into this category, where the period is a day (in this case "
today") and the metric is the number of visitors.

Here's how to write this query:

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// Build query
	todaysVisitorsQuery := plausible.AggregateQuery{
		Period: plausible.DayPeriod(),
		Metrics: plausible.Metrics{
			plausible.Visitors,
		},
	}

	// Make query
	result, err := mysite.Aggregate(todaysVisitorsQuery)
	if err != nil {
		// handle error
	}

	fmt.Printf("Total visitors of %s today: %d\n", mysite.ID(), result.Visitors)
}
```

### <a name="timeseries-queries"></a> Time Series Queries

A time series query reports a list of data points over a period of time, where each data point contains data about
metrics for that period of time.

A query like "the number of visitors and page views for each day in the 7 days before the 1st of February 2021"
falls into this category.

This is how to write this query:

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")
	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// Build query
	tsQuery := plausible.TimeseriesQuery{
		Period: plausible.Last7Days().FromDate(plausible.Date{Day: 1, Month: 2, Year: 2021}),
		Metrics: plausible.Metrics{
			plausible.Visitors,
			plausible.PageViews,
		},
	}

	// Make query
	queryResults, err := mysite.Timeseries(tsQuery)
	if err != nil {
		// handle error
	}

	// Iterate over the data points
	for _, stat := range queryResults {
		fmt.Printf("Date: %s | Visitors: %d | Pageviews: %d\n",
			stat.Date, stat.Visitors, stat.Pageviews)
	}

}
```

### <a name="breakdown-queries"></a> Breakdown Queries

A breakdown query reports stats for the value of a given property over a period of time.

For instance, a query like "over the last 7 days what are the number of visitors and page views for each page of my
site" falls into this category.

Here's how to write such query:

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	// Build query
	pageBreakdownQuery := plausible.BreakdownQuery{
		Property: plausible.EventPage,
		Period:   plausible.Last7Days(),
		Metrics: plausible.Metrics{
			plausible.Visitors,
			plausible.PageViews,
		},
	}

	// Make query
	pageBreakdown, err := mysite.Breakdown(pageBreakdownQuery)
	if err != nil {
		// handle error
	}

	// Iterate the results
	for _, stat := range pageBreakdown {
		fmt.Printf("Page: %s | Visitors: %d | Pageviews: %d \n",
			stat.Page, stat.Visitors, stat.Pageviews)
	}

}
```

## <a name="site-provisioning-api"></a> Site Provisioning API

This wrapper has support for the [site provisioning API](https://plausible.io/docs/sites-api).

However, note that this API is still private and requires a special token for the requests mentioned below to work. Make
sure you have a token with permissions for the site provisioning API before attempting to make these requests. You can
go here to know more about how to get a token for this API:

* [Plausible Docs: Site Provisioning API](https://plausible.io/docs/sites-api)

### <a name="provisioning-api-get-sites"></a> List sites

Gets a list of existing sites your Plausible account can access.

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	// Warning: This token must have permissions to the site provisioning API
	client := plausible.NewClient("<your_api_token>")

	sites, err := client.ListSites(plausible.ListSitesRequest{})

	if err != nil {
		// handle error
	}

	fmt.Printf("Sites %s\n", sites.Sites)
}
```

### <a name="provisioning-api-get-site"></a> Get site

Gets details of a site. Your Plausible account must have access to it.

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
	// Create a client with an API token
	// Warning: This token must have permissions to the site provisioning API
	client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")

	siteResult, err := mysite.Get()

	if err != nil {
		// handle error
	}

	fmt.Printf("Site %s\n", siteResult.Timezone)
}
```

### <a name="provisioning-api-shared-links"></a> Get or create Shared Links

Shared Links are URLs that you can generate to give others access to your dashboards.

You can use `SharedLink()` to get information for a link or create one with a given name. The call to get and create a
shared link it's the same - if a link with the given name already exists, it'll simply get the information for the
existent link. If the link does not exist, this call will create it and return the information of the newly created
link.

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
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
```

### <a name="provisioning-api-create-new-sites"></a> Create new sites

It's also possible to create new sites using the site provisioning API. Attempting to create a site that already exists
will result in an error.

```go
package main

import (
	"fmt"
	"github.com/andrerfcsantos/go-plausible/plausible"
)

func main() {
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
```

## <a name="events-api"></a> Events API

Push events with `site.PushEvent()`

```go
func main() {
    // Create a client with an API token
    client := plausible.NewClient("<your_api_token>")

	// Get an handler to perform queries for a given site
	mysite := client.Site("example.com")
	
      e := plausible.EventRequest {
		  EventData: EventData{
            Domain:  "example.org",
            Name:    "pageview",
            URL:     "https://example.com/awesome_page",
		  }
		  UserAgent:  "user-agent"
      }

      _, err := client.PushEvent(e)
      if err != nil {
        // handle error
      }
	}
}
```

## <a name="tests"></a> Tests

This project has tests in the form of Unit tests and Integration tests.

### <a name="unit-tests"></a> Unit Tests

Unit tests are the easiest to run as they don't require any setup and do not attempt to make requests over the internet.

Unit tests start with `TestUnit`. This means that to run just the unit tests, you can do:

```bash
go test github.com/andrerfcsantos/go-plausible/plausible -run ^TestUnit
```

### <a name="integration-tests"></a> Integration Tests

Integration tests attempt to make calls to the API. Because of this, they require configuration in the form of
environment variables. Set these environment variables before attempting to run the integration tests:

* `PLAUSIBLE_TOKEN` - API token to be used in the integration tests
* `PLAUSIBLE_DOMAINS` - A domain or a comma separated list of domains. The first domain on the list will be used to test
  queries.

Integration tests start with the name `TestIntegration`. With these variable set, you can run only the integration tests
with:

```bash
go test github.com/andrerfcsantos/go-plausible/plausible -run ^TestIntegration
```

To run the unit tests, and the integration tests, just omit the `-run` flag:

```bash
go test github.com/andrerfcsantos/go-plausible/plausible
```

These integration tests do not include tests that require the site provisioning API. See below you to active tests for
the site provisioning API.

### <a name="integration-tests-provisioning"></a> Integration Tests with the provisioning API

Integration tests to the site provisioning API are disabled by default.

There are a couple of reasons for this:

* The provisioning API is still private and requires a token with special permissions. Most users will use a regular API
  token, so these tests will not be relevant to them.

* The provisioning API allows the creation of sites and shared links, but the only way to reverse the actions of the API
  is by manually deleting them via the dashboard. This also means that **the cleanup for these tests must be done
  manually**.

With that said, if you really need to run these tests, set the following environment variable in addition to `PLAUSIBLE_TOKEN`
and `PLAUSIBLE_DOMAINS`:

* `PLAUSIBLE_PROVISIONING_TOKEN` - this must be set to an API token with permissions to the provisioning API.

With this variable set up, to run all tests (unit+integration tests) including the integration tests of the provisioning
API, add the flag `provisioning` to the `go test command`:

```bash
go test github.com/andrerfcsantos/go-plausible/plausible -flags=provisioning
```

## <a name="bugs-feedback"></a> Bugs and Feedback

If you encounter any bugs or have any comment or suggestion, please post them in
the [Issues section](https://github.com/andrerfcsantos/go-plausible/issues) of this repository.

## <a name="contributing"></a> Contributing

All contributions are welcome!

Feel free to open PR's or post suggestions on
the [Issues section](https://github.com/andrerfcsantos/go-plausible/issues).

## <a name="license"></a> License

This project uses the [MIT License](https://github.com/andrerfcsantos/go-plausible/blob/main/LICENSE)
