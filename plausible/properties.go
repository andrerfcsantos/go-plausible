package plausible

import "fmt"

// PropertyName represents the name of a property. Check the constants of this package to see a list of PropertyName
// values ready to use.
type PropertyName string

// PropertyName values:
const (
	// EventName is the name of the event name property
	EventName = PropertyName("event:name")
	// EventPage is the name of the event page property
	EventPage = PropertyName("event:page")
	// VisitSource is the name of the source property of a visit
	VisitSource = PropertyName("visit:source")
	// VisitReferrer is the name of the referrer property of a visit
	VisitReferrer = PropertyName("visit:referrer")
	// VisitUtmMedium is the name of utm medium property of a visit
	VisitUtmMedium = PropertyName("visit:utm_medium")
	// VisitUtmSource is the name of utm source property of a visit
	VisitUtmSource = PropertyName("visit:utm_source")
	// VisitUtmCampaign is the name of the utm campaign property of a visit
	VisitUtmCampaign = PropertyName("visit:utm_campaign")
	// VisitDevice is the name of the device property of a visit
	VisitDevice = PropertyName("visit:device")
	// VisitBrowser is the name of the browser property of a visit
	VisitBrowser = PropertyName("visit:browser")
	// VisitBrowserVersion is the name of the browser version property of a visit
	VisitBrowserVersion = PropertyName("visit:browser_version")
	// VisitOs is the name of the operating system property of a visit
	VisitOs = PropertyName("visit:os")
	// VisitOsVersion is the name of the operating system version property of a visit
	VisitOsVersion = PropertyName("visit:os_version")
	// VisitCountry is the name of the country property of a visit
	VisitCountry = PropertyName("visit:country")
)

// IsEmpty tells whether the name of the property is empty
func (pn *PropertyName) IsEmpty() bool {
	return string(*pn) == ""
}

func (pn *PropertyName) toQueryArgs() QueryArgs {
	return QueryArgs{
		{Name: "property", Value: string(*pn)},
	}
}

// Property represents a Plausible property, consisting of a name and value.
// Properties are used when building filters for queries
type Property struct {
	// Name is the name of the property
	Name PropertyName
	// Value is the value associated with the property
	Value string
}

func (p *Property) toFilterString() string {
	return fmt.Sprintf("%s==%s", p.Name, p.Value)
}

// CustomPropertyName makes a PropertyName for a custom property with a given name.
// A custom PropertyName is needed to create a Property related to a custom event.
// Also check the function CustomProperty for an easy way to create a property with
// a value and a custom name.
func CustomPropertyName(propertyName string) PropertyName {
	return PropertyName("event:props:" + propertyName)
}

// CustomProperty makes a Property out of a custom name and value for that property.
func CustomProperty(propertyName string, value string) Property {
	return Property{
		Name:  CustomPropertyName(propertyName),
		Value: value,
	}
}

// Properties represents a list of properties.
type Properties []Property

// Add adds a property to the list.
func (p *Properties) Add(property Property) {
	*p = append(*p, property)
}

// Count returns the number of properties in the list.
func (p *Properties) Count() int {
	return len(*p)
}
