package plausible

import "testing"

func TestUnitQueryArgsEquality(t *testing.T) {
	tests := []struct {
		name       string
		queryArgsA QueryArgs
		queryArgsB QueryArgs
		areEqual   bool
	}{
		{
			name: "equal query args",
			queryArgsA: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
			},
			queryArgsB: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
			},
			areEqual: true,
		},
		{
			name: "unequal query args with different lengths",
			queryArgsA: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
			},
			queryArgsB: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
				QueryArg{
					Name:  "field2",
					Value: "value2",
				},
			},
			areEqual: false,
		},
		{
			name: "unequal query args with same size but different names for query args",
			queryArgsA: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
			},
			queryArgsB: QueryArgs{
				QueryArg{
					Name:  "otherfield",
					Value: "value1",
				},
			},
			areEqual: false,
		},

		{
			name: "unequal query args with same size but different values for query args",
			queryArgsA: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "value1",
				},
			},
			queryArgsB: QueryArgs{
				QueryArg{
					Name:  "field1",
					Value: "othervalue",
				},
			},
			areEqual: false,
		},
	}

	for _, test := range tests {
		equal := test.queryArgsA.equalTo(test.queryArgsB)
		if test.areEqual != equal {
			t.Fatalf("test '%s' failed: unexpected result for equality, expected %v, got %v",
				test.name, test.areEqual, equal)
		}
	}

}
