package plausible

import "testing"

func TestUnitFilterString(t *testing.T) {

	tests := []struct {
		name                  string
		filter                Filter
		expectedFilterString  string
		expectedQueryArgs     QueryArgs
		expectedPropertyCount int
		isEmpty               bool
	}{
		{
			name:                  "empty filter",
			filter:                NewFilter(),
			expectedFilterString:  "",
			expectedQueryArgs:     QueryArgs{},
			expectedPropertyCount: 0,
			isEmpty:               true,
		},
		{
			name:                 "filter with 1 property created using NewFilter",
			filter:               NewFilter(Property{Name: EventName, Value: "pageview"}),
			expectedFilterString: "event:name==pageview",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "event:name==pageview"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with 2 properties created using NewFilter",
			filter:               NewFilter(Property{Name: EventName, Value: "pageview"}, Property{Name: VisitOs, Value: "Windows"}),
			expectedFilterString: "event:name==pageview;visit:os==Windows",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "event:name==pageview;visit:os==Windows"},
			},
			expectedPropertyCount: 2,
			isEmpty:               false,
		},
		{
			name:                 "filter with event name property",
			filter:               NewFilter().ByEventName("pageview"),
			expectedFilterString: "event:name==pageview",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "event:name==pageview"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with event page property",
			filter:               NewFilter().ByEventPage("/"),
			expectedFilterString: "event:page==/",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "event:page==/"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit source property",
			filter:               NewFilter().ByVisitSource("Google"),
			expectedFilterString: "visit:source==Google",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:source==Google"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit referrer property",
			filter:               NewFilter().ByVisitReferrer("example.com"),
			expectedFilterString: "visit:referrer==example.com",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:referrer==example.com"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit utm medium property",
			filter:               NewFilter().ByVisitUtmMedium("social"),
			expectedFilterString: "visit:utm_medium==social",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:utm_medium==social"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit utm source property",
			filter:               NewFilter().ByVisitUtmSource("twitter"),
			expectedFilterString: "visit:utm_source==twitter",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:utm_source==twitter"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit utm campaign property",
			filter:               NewFilter().ByVisitUtmCampaign("profile"),
			expectedFilterString: "visit:utm_campaign==profile",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:utm_campaign==profile"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit device property",
			filter:               NewFilter().ByVisitDevice("Desktop"),
			expectedFilterString: "visit:device==Desktop",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:device==Desktop"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit browser property",
			filter:               NewFilter().ByVisitBrowser("Firefox"),
			expectedFilterString: "visit:browser==Firefox",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:browser==Firefox"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit browser version property",
			filter:               NewFilter().ByVisitBrowserVersion("88.0.4324.146"),
			expectedFilterString: "visit:browser_version==88.0.4324.146",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:browser_version==88.0.4324.146"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit os property",
			filter:               NewFilter().ByVisitOs("Windows"),
			expectedFilterString: "visit:os==Windows",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:os==Windows"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit os version property",
			filter:               NewFilter().ByVisitOsVersion("10.6"),
			expectedFilterString: "visit:os_version==10.6",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:os_version==10.6"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with visit country property",
			filter:               NewFilter().ByVisitCountry("PT"),
			expectedFilterString: "visit:country==PT",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:country==PT"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
		{
			name:                 "filter with 2 properties created using builder pattern",
			filter:               NewFilter().ByVisitBrowser("Firefox").ByVisitCountry("PT"),
			expectedFilterString: "visit:browser==Firefox;visit:country==PT",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "visit:browser==Firefox;visit:country==PT"},
			},
			expectedPropertyCount: 2,
			isEmpty:               false,
		},
		{
			name:                 "filter with one custom property",
			filter:               NewFilter().ByCustomProperty("myproperty", "click"),
			expectedFilterString: "event:props:myproperty==click",
			expectedQueryArgs: QueryArgs{
				QueryArg{Name: "filters", Value: "event:props:myproperty==click"},
			},
			expectedPropertyCount: 1,
			isEmpty:               false,
		},
	}

	for _, test := range tests {
		actualFilterString := test.filter.toFilterString()
		if actualFilterString != test.expectedFilterString {
			t.Fatalf("test '%s' failed: expected filter string to be %s, actual filter string %s",
				test.name, test.expectedFilterString, actualFilterString)
		}

		actualQueryArgs := test.filter.toQueryArgs()
		if !actualQueryArgs.equalTo(test.expectedQueryArgs) {
			t.Fatalf("test '%s' failed: query args differ, expected %#v, got %#v",
				test.name, test.expectedQueryArgs, actualQueryArgs)
		}

		actualPropertyCount := test.filter.Count()
		if actualPropertyCount != test.expectedPropertyCount {
			t.Fatalf("test '%s' failed: expected %d properties for filter, got %d",
				test.name, test.expectedPropertyCount, actualPropertyCount)
		}

		actualEmptiness := test.filter.IsEmpty()
		if actualEmptiness != test.isEmpty {
			t.Fatalf("test '%s' failed: emptiness test failed: expected %v, got %v",
				test.name, test.isEmpty, actualEmptiness)
		}
	}

}
