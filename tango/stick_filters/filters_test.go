package stick_filters

import (
	"bytes"
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

func TestRawUrlEncode(t *testing.T) {
	filters := CreateFilters()
	filter := filters["rawurlencode"]

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
			`foo%20%22bar%22%20foo`},
		{
			"with unicode",
			func() stick.Value { return filter(nil, `foo 😀`) },
			`foo%20%F0%9F%98%80`,
		},
		{
			"with backslash",
			func() stick.Value { return filter(nil, `foo \ bar`) },
			`foo%20%5C%20bar`},
		{
			"with backslash and quotes",
			func() stick.Value { return filter(nil, `foo \ "bar"`) },
			`foo%20%5C%20%22bar%22`,
		},
		{
			"with newline",
			func() stick.Value { return filter(nil, "foo \n bar") },
			`foo%20%0A%20bar`,
		},
	}

	for _, test := range tests {
		actual := test.actual()
		if actual != test.expected {
			t.Errorf("rawurlencode(%s) returned an unexpected result: %s", test.name, actual)
		}
	}
}

func TestJsonDecode(t *testing.T) {
	filters := CreateFilters()
	filter := filters["json_decode"]

	// Test with a object
	func() {
		actual := filter(nil, `{"foo":"bar", "foobar": 3.1}`).(map[string]stick.Value)
		expected := map[string]stick.Value{"foo": "bar", "foobar": 3.1}
		// compare the values of the two maps
		for key, value1 := range expected {
			if value2, ok := actual[key]; !ok || value1 != value2 {
				t.Errorf("json_decode returned an unexpected result: %s", actual)
			}
		}
	}()

	// Test with a array
	func() {
		actual := filter(nil, `[1.0, 2.0, "foo", 3.0]`).([]stick.Value)
		expected := []stick.Value{1.0, 2.0, "foo", 3.0}
		// compare the values of the two maps
		for key, value1 := range expected {
			if value2 := actual[key]; value1 != value2 {
				t.Errorf("json_decode returned an unexpected result: %s", actual)
			}
		}
	}()

	// Test with array with objects
	func() {
		actual := filter(nil, `[{"foo": "bar"}, {"bar":"foo"}]`).([]stick.Value)
		expected := []stick.Value{map[string]stick.Value{"foo": "bar"}, map[string]stick.Value{"bar": "foo"}}
		// compare the values of the two maps
		for key, value1 := range expected {
			value2 := actual[key]
			// compare the map of value 1 to value 2
			for key, valueVal1 := range value1.(map[string]stick.Value) {
				if valueVal2, ok := value2.(map[string]stick.Value)[key]; !ok || valueVal1 != valueVal2 {
					t.Errorf(
						"json_decode returned an unexpected result: %s versus expected %s",
						valueVal2,
						valueVal1,
					)
				}
			}
		}
	}()

}

func TestBoolifyFilter(t *testing.T) {
	filters := CreateFilters()
	filter := filters["boolify"]

	trueValues := []interface{}{
		stick.Value("true"),
		stick.Value(1.1),
		stick.Value(true),
		true,
		1,
	}
	falseValues := []interface{}{
		stick.Value("false"),
		stick.Value(0.0),
		stick.Value(false),
		false,
		0,
		nil,
	}
	//errorValues := []interface{}{"foo", []string{"foo"}, map[string]string{"foo": "bar"}}

	for _, trueValue := range trueValues {
		result := filter(nil, trueValue).(bool)
		if result != true {
			t.Errorf("boolify returned an unexpected result: %v", result)
		}
	}
	for _, falseValue := range falseValues {
		result := filter(nil, falseValue).(bool)
		if result != false {
			t.Errorf("boolify returned an unexpected result: %v", result)
		}
	}
}

func TestBoolSwitchFilter(t *testing.T) {
	filters := CreateFilters()
	filter := filters["bool_switch"]

	result := filter(nil, true, "ok", "not ok")
	if result != "ok" {
		t.Errorf("bool_switch returned an unexpected result: %v", result)
	}

	result = filter(nil, false, "ok", "not ok")
	if result != "not ok" {
		t.Errorf("bool_switch returned an unexpected result: %v", result)
	}

	result = filter(nil, "disabled", "ok", "not ok")
	if result != "not ok" {
		t.Errorf("bool_switch returned an unexpected result: %v", result)
	}
}

func TestExistsFilter(t *testing.T) {
	filters := CreateFilters()
	filter := filters["exists"]

	result := filter(nil, "foo")
	if result != false {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	result = filter(nil, nil)
	if result != false {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	result = filter(nil, "")
	if result != false {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	result = filter(nil, 0)
	if result != false {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	result = filter(nil, false)
	if result != false {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	loader := stick.StringLoader{}
	var st = stick.New(&loader)
	st.Filters = CreateFilters()

	buf := new(bytes.Buffer)

	st.Execute(
		"{% if \"foo\"|exists %}1{% else %}0{% endif %}",
		buf, map[string]stick.Value{"foo": "bar"},
	)
	result = buf.String()
	if result != "1" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	buf.Truncate(0)
	st.Execute(
		"{% if \"foo\"|exists %}1{% else %}0{% endif %}",
		buf,
		map[string]stick.Value{"bar": "foo"},
	)
	result = buf.String()
	if result != "0" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}
}

func TestValueFilter(t *testing.T) {
	loader := stick.StringLoader{}
	var st = stick.New(&loader)
	st.Filters = CreateFilters()

	buf := new(bytes.Buffer)
	st.Execute(
		"{{ \"foo\"|value }}",
		buf,
		map[string]stick.Value{"foo": "bar"},
	)
	result := buf.String()
	if result != "bar" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	buf.Truncate(0)
	st.Execute(
		"{{ \"foo\"|value(\"foo\") }}",
		buf,
		map[string]stick.Value{"foo": "bar"},
	)
	result = buf.String()
	if result != "bar" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	buf.Truncate(0)
	st.Execute(
		"{{ \"bar\"|value }}",
		buf,
		map[string]stick.Value{"foo": "bar"},
	)
	result = buf.String()
	if result != "" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}

	buf.Truncate(0)
	st.Execute(
		"{{ \"bar\"|value(\"default\") }}",
		buf,
		map[string]stick.Value{"foo": "bar"},
	)
	result = buf.String()
	if result != "default" {
		t.Errorf("exists returned an unexpected result: %v", result)
	}
}
