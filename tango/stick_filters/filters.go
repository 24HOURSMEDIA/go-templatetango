package stick_filters

import (
	"encoding/json"
	"github.com/tyler-sommer/stick"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func CreateFilters() map[string]stick.Filter {
	return map[string]stick.Filter{
		// encodes a value to JSON, i.e. {  "foo": {{ value|json_value }} }
		"json_value":        jsonValue,
		"json_casted_value": jsonCastedValue,
		"json_escape":       jsonEscape,
		"rawurlencode":      rawUrlEncode,
		"json_decode":       jsonDecode,
		"boolify":           boolifyFilter,
		"bool_switch":       boolSwitchFilter,
		"exists":            existsFilter,
		"value":             valueFilter,
		"apply_mapping":     applyMappingFilter,
		"fatality":          fatalityFilter,
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

func rawUrlEncode(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	str := stick.CoerceString(val)
	return url.PathEscape(str)
}

func jsonDecode(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	str := stick.CoerceString(val)
	var v stick.Value
	err := json.Unmarshal([]byte(str), &v)
	if err != nil {
		log.Fatal("jsonDecode error: ", err)
	}
	return stickify(v)
}

func boolifyFilter(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	v, err := boolify(val)
	if err != nil {
		log.Fatal("boolify error: ", err)
	}
	return stickify(v)
}

func boolSwitchFilter(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	if len(args) < 2 {
		log.Fatal("boolSwitch error: two arguments must be provided, a true value and a false value")
	}
	boolVal, err := boolify(val)
	if err != nil {
		log.Fatal("boolSwitch error: ", err)
	}
	if boolVal {
		return stickify(args[0])
	}
	return stickify(args[1])
}

// existsFilter checks if a variable exists in the current context
func existsFilter(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	if (ctx == nil) || (ctx.Scope() == nil) {
		return stickify(false)
	}
	var name = stick.CoerceString(val)
	_, exists := ctx.Scope().All()[name]
	return stickify(exists)
}

// valueFilter returns the value of a variable in the current scope and context,
// or a default value if it does not exist
func valueFilter(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	if (ctx == nil) || (ctx.Scope() == nil) {
		return stickify(nil)
	}
	var name = stick.CoerceString(val)
	defaultVal := stickify(nil)
	if len(args) > 0 {
		defaultVal = stickify(args[0])
	}
	value, exists := ctx.Scope().All()[name]
	if exists {
		return stickify(value)
	}
	return defaultVal
}
