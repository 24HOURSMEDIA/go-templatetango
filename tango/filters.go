package tango

import (
	"encoding/json"
	"github.com/tyler-sommer/stick"
)

func CreateFilters() map[string]stick.Filter {
	return map[string]stick.Filter{
		// encodes a value to JSON, i.e. {  "foo": {{ value|json_value }} }
		"json_value": jsonValue,
	}
}

func jsonValue(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	input := val
	encoded, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	// Convert the byte slice to a string and return
	return string(encoded)
}
