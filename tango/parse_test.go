package tango

import (
	"github.com/tyler-sommer/stick"
	"testing"
)

func TestParseString(t *testing.T) {
	template := "Hello, {{ name }}!"
	params := map[string]stick.Value{"name": "FooBar"}
	result, err := ParseString(template, &params)
	if err != nil {
		t.Errorf("Parse() returned an error: %s", err)
	}
	if result != "Hello, FooBar!" {
		t.Errorf("Parse() returned an unexpected result: %s", result)
	}
}

func TestParseStringWithEscaping(t *testing.T) {
	// Ensure that values in 'html' remain unescaped
	// It is the responsibility of the template author to ensure that

	template := "HTML: \"{{ html }}\""
	params := map[string]stick.Value{"html": "a&b"}
	result, err := ParseString(template, &params)
	if err != nil {
		t.Errorf("Parse() returned an error: %s", err)
	}
	if result != "HTML: \"a&b\"" {
		t.Errorf("Parse() returned an unexpected result: %s", result)
	}

	template = "<a href=\"{{ url }}\">{{ url }}</a>"
	params = map[string]stick.Value{"url": "https://example.com/?a=b&c=d"}
	result, err = ParseString(template, &params)
	if err != nil {
		t.Errorf("Parse() returned an error: %s", err)
	}
	if result != "<a href=\"https://example.com/?a=b&c=d\">https://example.com/?a=b&c=d</a>" {
		t.Errorf("Parse() returned an unexpected result: %s", result)
	}
}
