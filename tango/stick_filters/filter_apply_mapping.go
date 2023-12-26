package stick_filters

import (
	"github.com/tyler-sommer/stick"
)

// applyMappingWithSuffix remaps a map of keys to new keys, using the suffix to
// append to the key to look up the value in the source map. If the value
// is not found, the default value is used.
func applyMappingWithSuffix(
	mapping map[string]string,
	src map[string]stick.Value,
	suffix string,
	defaultVal stick.Value,
) *map[string]stick.Value {
	target := make(map[string]stick.Value)
	for key1, key2 := range mapping {
		val, exists := src[key2+suffix]
		if !exists {
			val = defaultVal
		}
		target[key1] = val
	}
	return &target
}

// remapFilter remaps a map of keys to new keys, using the suffix to
// append to the key to look up the value in the source map. If the value
// is not found, the default value is used.
func applyMappingFilter(
	ctx stick.Context,
	val stick.Value,
	args ...stick.Value,
) stick.Value {
	mapping, err := makeStringMap(val.(map[string]stick.Value))
	if err != nil {
		return nil
	}

	var src map[string]stick.Value
	if (len(args) > 0) && (args[0] != nil) {
		result, ok := args[0].(map[string]stick.Value)
		if !ok {
			return nil
		}
		src = result
	} else {
		src = ctx.Scope().All()
	}

	var suffix string
	if len(args) > 1 {
		suffix = stick.CoerceString(args[1])
	}

	var defaultVal interface{}
	if len(args) > 2 {
		defaultVal = args[2]
	}

	return stickify(applyMappingWithSuffix(mapping, src, suffix, defaultVal))
}
