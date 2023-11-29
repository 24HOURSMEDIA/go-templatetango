package tango

import (
	"github.com/tyler-sommer/stick"
	"testing"
)

func TestJsonValue(t *testing.T) {
	filters := CreateFilters()
	filter := filters["json_value"]

	tests := []struct {
		name     string
		actual   func() stick.Value
		expected stick.Value
	}{
		{"string", func() stick.Value { return filter(nil, "foo") }, `"foo"`},
		{"integer", func() stick.Value { return filter(nil, 123) }, `123`},
		{"float", func() stick.Value { return filter(nil, 123.2) }, `123.2`},
		{"bool-true", func() stick.Value { return filter(nil, true) }, `true`},
		{"bool-false", func() stick.Value { return filter(nil, false) }, `false`},
		{"null", func() stick.Value { return filter(nil, nil) }, `null`},
	}

	for _, test := range tests {
		actual := test.actual()
		if actual != test.expected {
			t.Errorf("json_value(%s) returned an unexpected result: %s", test.name, actual)
		}
	}
}
