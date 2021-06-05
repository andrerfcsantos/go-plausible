/*
Package plausible implements a client/wrapper for the Plausible Analytics API.

It currently supports the Stats API and the Site Provisioning API.

Making a client and site handlers

Start by creating a client to make requests. Then, request handlers to specific sites from the client:

    // Create a client with an API token
    client := plausible.NewClient("<your_api_token>")

    // Get an handler to perform queries for a given site
    mysite := client.Site("example.com")

    // You can reuse the same client to get handlers for additional sites
    myothersite := client.Site("otherexample.com")

    // Use 'mysite' and the 'myothersite' handlers to query stats for the sites

Queries - Stats API

You can use the site handlers to perform queries / requests for data. There are four types of queries: current visitor queries,
aggregate queries, time series queries and breakdown queries. This queries can be done using the methods on a site handlers
called CurrentVisitors, Aggregate, Timeseries and Breakdown respectively. Check the documentation of each one of these methods
for query examples.

The current visitors query is the most straight forward one - for a given site return the number of current visitors.
Check the method CurrentVisitors for more information and for examples of this query.

An aggregate query reports data for metrics aggregated over a period of time. A query like "the total number of visitors
today" fall into this category, where the period is a day (in this case "today") and the metric is the number of visitors.
Check the method Aggregate for more information and for examples of this query.

A time series query reports a list of data points over a period of time, where each data point contains data about
metrics for that period of time. A query like "the number of visitors and page views for each day in the 7 days before
the 1st of February 2021" falls into this category.
Check the method Timeseries for more information and for examples of this query.

A breakdown query reports stats for the value of a given property over a period of time. For instance, a query like
"over the last 7 days what are the number of visitors and page views for each page of my site" falls into this category.
Check the method Breakdown for more information and for examples of this query.

Provisioning API

The provisioning API allows to create new sites on Plausible and to create shared links for those sites.
The methods CreateNewSite and SharedLink respectively implement these requests.

However, please note that these methods are using the provisioning API which requires a token with special permissions
for the requests to succeed. For more info: https://plausible.io/docs/sites-api

Time periods

When requesting aggregate information to the API, it's only possible to get data for a given period of time, for instance,
"the last 7 days", "the last month" or "the current day" are examples of time periods.
All time periods are relative to a date. When the date information is missing from a time period,
the date is assumed to be "today". It's also possible to specify a time period between two specific dates.

Time periods are represented by the TimePeriod type.

Properties

Each pageview or custom event has some properties associated with it. These properties can be
used when querying the API to filter the results.

Properties are represented by the Property type.
Properties have a name and value. The name of a property is represented by the PropertyName type.
Typically, most users won't need to make custom property names and can just use the constant PropertyName
values declared at the top-level of the package like VisitOs or VisitBrowser.

Filters

Filters allow to drill down and segment the data the results refer to. All queries for data accept an optional
filter argument.

Filters are represented by the Filter type is consists of a simple list of properties by
which you want to filter.

Metrics

Metrics are aggregate information about the data. All queries have the option for you to choose the metrics
you want to see included in the results.

There are four metrics currently that you can ask the results for: number of visitors, number of page views,
visit duration and bounce rate.
These metrics are represented by the Metric type and there are four constants of this type,
each one representing one of the four metrics: Visitors, PageViews, BounceRate and VisitDuration.

Time Intervals

Time intervals are used for time series queries to specify the interval of time between two consecutive data points.

A time interval is represented by the TimeInterval type.
There are currently two time intervals supported: date and month which are represented with the DateInterval
and MonthInterval constants respectively.

A MonthInterval means a month of difference between data points. For instance, if you ask for time series data
over the last 6 months with a month interval, this means you will get 6 data points back - 1 for each month.

A DateInterval, depending on the query, means a day or an hour of difference between each data point. For instance, if
you ask for time series data over the last 30 days with a date interval, you will get 30 data points back - 1 for each day.
However, with a DateInterval, when the period of the time series refers to a day, for instance "today", the data points will
actually have 1 hour of interval between them. You can check the Date string field of each data point to know
about which date/hour the data refers to.

Useful Links

Stats API: https://plausible.io/docs/stats-api

Site provisioning API: https://plausible.io/docs/sites-api

Concepts: https://plausible.io/docs/stats-api#concepts

*/
package plausible
