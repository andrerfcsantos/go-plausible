package plausible

// Filter represents an API rawAggregateResult filter over properties of the stats data.
// The filter is a logic AND over all its properties and values.
// Filters are used when making a request to the API to narrow the information returned.
type Filter struct {
	// Properties to filter by
	Properties Properties
}

// NewFilter creates a new filter with the given properties.
func NewFilter(properties ...Property) Filter {

	f := Filter{}

	for _, property := range properties {
		f.Properties.Add(property)
	}

	return f
}

// ByEventName adds a filter over the event name property to the current filter.
func (f Filter) ByEventName(eventName string) Filter {
	f.Properties.Add(Property{Name: EventName, Value: eventName})
	return f
}

// ByEventPage adds a filter over the page property to the current filter.
func (f Filter) ByEventPage(page string) Filter {
	f.Properties.Add(Property{Name: EventPage, Value: page})
	return f
}

// ByVisitSource adds a filter over the source property to the current filter.
func (f Filter) ByVisitSource(source string) Filter {
	f.Properties.Add(Property{Name: VisitSource, Value: source})
	return f
}

// ByVisitReferrer adds a filter over the referrer property to the current filter.
func (f Filter) ByVisitReferrer(referrer string) Filter {
	f.Properties.Add(Property{Name: VisitReferrer, Value: referrer})
	return f
}

// ByVisitUtmMedium adds a filter over the utm medium property to the current filter.
func (f Filter) ByVisitUtmMedium(utmMedium string) Filter {
	f.Properties.Add(Property{Name: VisitUtmMedium, Value: utmMedium})
	return f
}

// ByVisitUtmSource adds a filter over the utm source property to the current filter.
func (f Filter) ByVisitUtmSource(utmSource string) Filter {
	f.Properties.Add(Property{Name: VisitUtmSource, Value: utmSource})
	return f
}

// ByVisitUtmCampaign adds a filter over the utm campaign property to the current filter.
func (f Filter) ByVisitUtmCampaign(utmCampaign string) Filter {
	f.Properties.Add(Property{Name: VisitUtmCampaign, Value: utmCampaign})
	return f
}

// ByVisitDevice adds a filter over the device property to the current filter.
func (f Filter) ByVisitDevice(device string) Filter {
	f.Properties.Add(Property{Name: VisitDevice, Value: device})
	return f
}

// ByVisitBrowser adds a filter over the browser property to the current filter.
func (f Filter) ByVisitBrowser(browser string) Filter {
	f.Properties.Add(Property{Name: VisitBrowser, Value: browser})
	return f
}

// ByVisitBrowserVersion adds a filter over the browser version property to the current filter.
func (f Filter) ByVisitBrowserVersion(browserVersion string) Filter {
	f.Properties.Add(Property{Name: VisitBrowserVersion, Value: browserVersion})
	return f
}

// ByVisitOs adds a filter over the operating system property to the current filter.
func (f Filter) ByVisitOs(operatingSystem string) Filter {
	f.Properties.Add(Property{Name: VisitOs, Value: operatingSystem})
	return f
}

// ByVisitOsVersion adds a filter over the operating system version property to the current filter.
func (f Filter) ByVisitOsVersion(osVersion string) Filter {
	f.Properties.Add(Property{Name: VisitOsVersion, Value: osVersion})
	return f
}

// ByVisitCountry adds a filter over the country property to the current filter.
func (f Filter) ByVisitCountry(country string) Filter {
	f.Properties.Add(Property{Name: VisitCountry, Value: country})
	return f
}

// ByCustomProperty adds a filter over a custom property to the current filter.
func (f Filter) ByCustomProperty(propertyName string, value string) Filter {
	f.Properties.Add(Property{Name: CustomPropertyName(propertyName), Value: value})
	return f
}

// Count returns the number of properties in the filter
func (f Filter) Count() int {
	return len(f.Properties)
}

func (f Filter) toFilterString() string {

	s := ""
	n := f.Properties.Count()

	for i := 0; i < n; i++ {
		s += f.Properties[i].toFilterString()

		if i != n-1 {
			s += ";"
		}
	}

	return s
}

func (f Filter) toQueryArgs() QueryArgs {

	var qArgs QueryArgs

	if !f.IsEmpty() {
		qArgs = append(qArgs, QueryArg{Name: "filters", Value: f.toFilterString()})
	}

	return qArgs
}

// IsEmpty tells if the filter has no properties
func (f Filter) IsEmpty() bool {
	return f.Properties.Count() == 0
}
