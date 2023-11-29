package tango

import (
	"encoding/json"
	"github.com/tyler-sommer/stick"
	"strconv"
	"strings"
)

func CreateFilters() map[string]stick.Filter {
	return map[string]stick.Filter{
		// encodes a value to JSON, i.e. {  "foo": {{ value|json_value }} }
		"json_value":        jsonValue,
		"json_casted_value": jsonCastedValue,
		"json_escape":       jsonEscape,
	}
}

// jsonValue encodes a value to JSON, i.e. {  "foo": {{ value|json_value }} }
func jsonValue(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	input := val
	encoded, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	// Convert the byte slice to a string and return
	return string(encoded)
}

// jsonCastValue encodes a value to JSON, i.e. {  "foo": {{ value|json_value }} }
func jsonCastedValue(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	str := stick.CoerceString(val)
	normalizedStr := strings.TrimSpace(strings.ToLower(str))
	if canCastToNumber(str) {
		v, _ := strconv.ParseFloat(str, 64)
		return jsonValue(ctx, v, args...)
	}
	// if normalized string is "true", "false"
	if normalizedStr == "true" || normalizedStr == "false" {
		v, _ := strconv.ParseBool(normalizedStr)
		return jsonValue(ctx, v, args...)
	}
	// if normalized string is "null""
	if normalizedStr == "null" {
		return jsonValue(ctx, nil, args...)
	}

	return jsonValue(ctx, str, args...)
}

func jsonEscape(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	str := stick.CoerceString(val)
	encoded, err := json.Marshal(str)
	if err != nil {
		return ""
	}
	// From encoded, strip the first and last character:
	return string(encoded[1 : len(encoded)-1])
}

// isFloat checks if the given string is a valid floating point number.
func canCastToNumber(s string) bool {
	_, err := strconv.ParseFloat(s, 64) // 64 refers to the precision in bits
	return err == nil
}
