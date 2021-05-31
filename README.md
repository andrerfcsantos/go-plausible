# `go-plausible` - Go Wrapper for the Plausible API

[![Go Reference](https://pkg.go.dev/badge/github.com/andrerfcsantos/go-plausible/plausible.svg)](https://pkg.go.dev/github.com/andrerfcsantos/go-plausible/plausible) [![Go Report Card](https://goreportcard.com/badge/github.com/andrerfcsantos/go-plausible)](https://goreportcard.com/report/github.com/andrerfcsantos/go-plausible)

Go wrapper/client for the [Plausible](https://plausible.io/) API.

This wrapper currently supports the [Stats API](https://plausible.io/docs/stats-api).

Support for the currently private [Site Provisioning API](https://plausible.io/docs/sites-api) is planned.

## Usage

```go
import "github.com/andrerfcsantos/go-plausible/plausible"
```

To use this client, you'll need an API token, which you can get from Plausible.

With the API token, create a client using the token and get an handler for one or more sites:
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

    // Use 'mysite' and the 'myothersite' handlers to query stats for the sites
    // ...
}
```

## Queries

There are 4 types of queries supported by the API:

* Current Visitors
* Aggregate Queries
* Timeseries Queries
* Breakdown Queries

### Current Visitors

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
    fmt.Printf("Site %s has %d current visitors!\n", err.Error())
}
```

### Aggregate Queries

An aggregate query reports data for metrics aggregated over a period of time.

A query like "the total number of visitors today" fall into this category,
where the period is a day (in this case "today") and the metric is the number of visitors.

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
    fmt.Printf("Total visitors of %s today: %d\n", site.ID(), result.Visitors)
}
```

### Time Series Queries

A time series query reports a list of data points over a period of time,
where each data point contains data about metrics for that period of time.

A query like "the number of visitors and page views for each day in the 7 days before the 1st of February 2021" falls into this category.

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

	tsQuery := plausible.TimeseriesQuery{
		Period:  plausible.Last7Days().FromDate(plausible.Date{1, 2, 2021}),
		Metrics: plausible.Metrics {
			plausible.Visitors,
			plausible.PageViews,
		},
	}


	queryResults, err := site.Timeseries(tsQuery)
	if err != nil {
		// handle error
	}

	// Iterate over the data points
	for _, stat := range queryResults {
    	fmt.Printf("\tDate: %s | Visitors: %d | Pageviews: %d\n",
    		    stat.Date, stat.Visitors, stat.Pageviews)
    }

}
```


### Breakdown Queries

A breakdown query reports stats for the value of a given property over a period of time.

For instance, a query like "over the last 7 days what are the number of visitors and
page views for each page of my site" falls into this category.

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

    pageBreakdownQuery := plausible.BreakdownQuery{
    	Property: plausible.EventPage,
    	Period:   plausible.Last7Days(),
    	Metrics:  plausible.Metrics {
    		plausible.Visitors,
    		plausible.PageViews,
    	},
    }

    pageBreakdown, err := site.Breakdown(pageBreakdownQuery)
    if err != nil {
    	// handle error
    }

    for _, stat := range pageBreakdown {
    	fmt.Printf("Page: %s | Visitors: %d | Pageviews: %d \n",
    	            stat.Page, stat.Visitors, stat.Pageviews)
    }

}
```

## Bugs and Feedback

If you encounter any bugs or have any comment or suggestion, please post them in the [Issues section](https://github.com/andrerfcsantos/go-plausible/issues) of this repository.

## Contributing

All contributions are welcome!

Feel free to open PR's or post suggestions on the [Issues section](https://github.com/andrerfcsantos/go-plausible/issues).

## License

This project uses the [MIT License](https://github.com/andrerfcsantos/go-plausible/blob/main/LICENSE)
