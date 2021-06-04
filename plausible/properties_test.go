package plausible

import "testing"

func TestUnitIndividualProperties(t *testing.T) {
	tests := []struct {
		name                 string
		property             Property
		expectedName         PropertyName
		expectedValue        string
		expectedFilterString string
	}{
		{
			name:                 "custom property test",
			property:             CustomProperty("myproperty", "profile"),
			expectedName:         PropertyName("event:props:myproperty"),
			expectedValue:        "profile",
			expectedFilterString: "event:props:myproperty==profile",
		},
	}

	for _, test := range tests {

		if test.expectedName != test.property.Name {
			t.Fatalf("test '%s' failed: properties have different names %v != %v",
				test.name, test.expectedName, test.property.Name)
		}

		if test.expectedValue != test.property.Value {
			t.Fatalf("test '%s' failed: properties have the same name (%v) but different values %v != %v",
				test.name, test.expectedName, test.expectedValue, test.property.Value)
		}
		actualFilterString := test.property.toFilterString()

		if test.expectedFilterString != actualFilterString {
			t.Fatalf("test '%s' failed: expected filter string %s, got %s",
				test.name, test.expectedFilterString, actualFilterString)
		}

	}
}
