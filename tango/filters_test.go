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

func TestJsonCastValue(t *testing.T) {
	filters := CreateFilters()
	filter := filters["json_casted_value"]

	tests := []struct {
		name     string
		actual   func() stick.Value
		expected stick.Value
	}{
		{"string", func() stick.Value { return filter(nil, "foo") }, `"foo"`},
		{"integer", func() stick.Value { return filter(nil, "123") }, `123`},
		{"float", func() stick.Value { return filter(nil, "123.2") }, `123.2`},
		{"bool-true", func() stick.Value { return filter(nil, "true") }, `true`},
		{"bool-false", func() stick.Value { return filter(nil, "false") }, `false`},
		{"bool-true-ucase", func() stick.Value { return filter(nil, "TRUE") }, `true`},
		{"bool-false-ucase", func() stick.Value { return filter(nil, "False") }, `false`},
		{"null", func() stick.Value { return filter(nil, "null") }, `null`},
		{"null-ucase", func() stick.Value { return filter(nil, "Null") }, `null`},
	}

	for _, test := range tests {
		actual := test.actual()
		if actual != test.expected {
			t.Errorf("json_casted_value(%s) returned an unexpected result: %s", test.name, actual)
		}
	}
}

func TestJsonEscape(t *testing.T) {
	filters := CreateFilters()
	filter := filters["json_escape"]

	tests := []struct {
		name     string
		actual   func() stick.Value
		expected stick.Value
	}{
		{
			"normal",
			func() stick.Value { return filter(nil, "foo") },
			`foo`},
		{
			"with quotes",
			func() stick.Value { return filter(nil, `foo "bar" foo`) },
			`foo \"bar\" foo`},
		{
			"with unicode",
			func() stick.Value { return filter(nil, `foo 😀`) },
			`foo 😀`,
		},
		{
			"with backslash",
			func() stick.Value { return filter(nil, `foo \ bar`) },
			`foo \\ bar`},
		{
			"with backslash and quotes",
			func() stick.Value { return filter(nil, `foo \ "bar"`) },
			`foo \\ \"bar\"`,
		},
		{
			"with newline",
			func() stick.Value { return filter(nil, "foo \n bar") },
			`foo \n bar`,
		},
	}

	for _, test := range tests {
		actual := test.actual()
		if actual != test.expected {
			t.Errorf("json_str_encode(%s) returned an unexpected result: %s", test.name, actual)
		}
	}
}
