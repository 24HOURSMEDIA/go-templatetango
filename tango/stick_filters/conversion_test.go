package stick_filters

import (
	"github.com/tyler-sommer/stick"
	"testing"
)

func TestBoolify(t *testing.T) {
	trueValues := []interface{}{
		true,
		"True",
		"true",
		"1",
		"2",
		"-1",
		"1.1",
		"0.1",
		"ON",
		"on",
		"yes",
		"y",
		"EnAbLeD",
		"enabled",
		"enable",
		1,
		2,
		0.1,
		stick.Value("true"),
		stick.Value(1.1),
		stick.Value(true),
	}
	falseValues := []interface{}{
		false,
		"false",
		"0",
		"0.0",
		"off",
		"OFF",
		"no",
		"No",
		"n",
		"disabled",
		"disable",
		0,
		0.0,
		stick.Value("false"),
		stick.Value(0.0),
		stick.Value(false),
	}
	errorValues := []interface{}{"foo", []string{"foo"}, map[string]string{"foo": "bar"}}

	for _, trueValue := range trueValues {
		result, err := boolify(trueValue)
		if err != nil {
			t.Errorf("boolify returned an error: %s", err)
		}
		if result != true {
			t.Errorf("boolify returned an unexpected result: %v", result)
		}
	}
	for _, falseValue := range falseValues {
		result, err := boolify(falseValue)
		if err != nil {
			t.Errorf("boolify returned an error: %s", err)
		}
		if result != false {
			t.Errorf("boolify returned an unexpected result: %v", result)
		}
	}
	for _, errorValue := range errorValues {
		result, err := boolify(errorValue)
		if result != false {
			t.Errorf("boolify(%v) returned an unexpected result for an error: %v", errorValue, result)
		}
		if err == nil {
			t.Errorf("boolify(%v) did not return an error", errorValue)
		}
	}
}
