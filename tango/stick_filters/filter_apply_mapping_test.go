package stick_filters

import (
	"bytes"
	"github.com/tyler-sommer/stick"
	"strings"
	"testing"
)

func TestApplyMappingWithSuffix(t *testing.T) {
	src := map[string]stick.Value{
		"bar":        "bar_value",
		"qux":        "qux_value",
		"bar_suffix": "bar_value_from_suffix",
		"qux_suffix": "qux_value_from_suffix",
	}
	func() {
		// Test with no suffix
		mapping := map[string]string{
			"attr1": "bar",
			"attr2": "qux",
		}
		suffix := ""
		defaultVal := "default"
		result := applyMappingWithSuffix(mapping, src, suffix, defaultVal)
		if len(*result) != 2 {
			t.Errorf("applyMappingWithSuffix returned an unexpected length: %v", result)
		}
		if (*result)["attr1"] != "bar_value" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
		if (*result)["attr2"] != "qux_value" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
	}()
	func() {
		// Test with suffix
		mapping := map[string]string{
			"attr1": "bar",
			"attr2": "qux",
		}
		suffix := "_suffix"
		defaultVal := "default"
		result := applyMappingWithSuffix(mapping, src, suffix, defaultVal)
		if len(*result) != 2 {
			t.Errorf("applyMappingWithSuffix returned an unexpected length: %v", result)
		}
		if (*result)["attr1"] != "bar_value_from_suffix" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
		if (*result)["attr2"] != "qux_value_from_suffix" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
	}()
	func() {
		// Test with explicit default value
		mapping := map[string]string{
			"attr1":         "bar",
			"attr2":         "qux",
			"attr_notexist": "notexist",
		}
		suffix := "_suffix"
		defaultVal := "default"
		result := applyMappingWithSuffix(mapping, src, suffix, defaultVal)
		if len(*result) != 3 {
			t.Errorf("applyMappingWithSuffix returned an unexpected length: %v", result)
		}
		if (*result)["attr1"] != "bar_value_from_suffix" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
		if (*result)["attr2"] != "qux_value_from_suffix" {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
		if (*result)["attr_notexist"] != defaultVal {
			t.Errorf("applyMappingWithSuffix returned an unexpected result: %v", result)
		}
	}()
}

func TestApplyMappingFilter(t *testing.T) {
	loader := stick.StringLoader{}
	var st = stick.New(&loader)
	st.Filters = map[string]stick.Filter{
		"apply_mapping": applyMappingFilter,
	}

	// create a multiline string
	multiline := `
{% set var1 = "var1val" %}
{% set var2 = "var2val" %}
{% set mapping = {"foo": "var1", "baz": "var2"} %}
{% set obj = mapping | apply_mapping %}
result:{{ obj.foo }}-{{ obj.baz }}
`
	buf := new(bytes.Buffer)
	st.Execute(
		multiline,
		buf,
		map[string]stick.Value{},
	)
	result := strings.Trim(buf.String(), "\n ")
	if result != "result:var1val-var2val" {
		t.Errorf("remap returned an unexpected result: %v", result)
	}

}
