/*
Package plausible implements a client/wrapper for the Plausible Analytics API.

Making a client

To start interacting with the API, a client must be created. After that, it's possible to work with multiple sites by requesting site handlers to the client:

    // Create a client with an API token
    client := plausible.NewClient("<your_api_token>")

    // Get an handler to perform queries for a given site
    mysite := client.Site("example.com")

    // You can reuse the same client to get handlers for additional sites
    myothersite := client.Site("otherexample.com")

    // Use 'mysite' and the 'myothersite' handlers to query stats for the sites

Current visitors query

This is the most straight forward query - for a given site return the number of current visitors:

    visitors, err := mysite.CurrentVisitors()
    if err != nil {
    	// handle error
    }

Aggregate queries

An aggregate query reports data for metrics aggregated over a period of time.

A query like "the total number of visitors today" fall into this category,
where the period is a day (in this case "today") and the metric is the number of visitors.

Here's how to write this query:

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

    // inspect result ...

Time Series queries

A time series query reports a list of data points over a period of time,
where each data point contains data about metrics for that period of time.

A query like "the number of visitors and page views for each day in the 7 days before the 1st of February 2021" falls into this category.

This is how to write this query:

	tsQuery := plausible.TimeseriesQuery{
		Period:  plausible.Last7Days().FromDate(plausible.Date{Day: 1, Month: 2, Year: 2021}),
		Metrics: plausible.Metrics {
			plausible.Visitors,
			plausible.PageViews,
		},
	}

	queryResults, err := mysite.Timeseries(tsQuery)
	if err != nil {
		// handle error
	}
	// Iterate over queryResults ...

Breakdown queries

A breakdown query reports stats for the value of a given property over a period of time.

For instance, a query like "over the last 7 days what are the number of visitors and
page views for each page of my site" falls into this category.

Here's how to write such query:

    pageBreakdownQuery := plausible.BreakdownQuery{
    	Property: plausible.EventPage,
    	Period:   plausible.Last7Days(),
    	Metrics:  plausible.Metrics {
    		plausible.Visitors,
    		plausible.PageViews,
    	},
    }

    pageBreakdown, err := mysite.Breakdown(pageBreakdownQuery)
    if err != nil {
    	// handle error
    }

    // Iterate pageBreakdown ...

*/
package plausible
