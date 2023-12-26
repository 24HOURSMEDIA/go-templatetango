package stick_filters

import "github.com/tyler-sommer/stick"

// makeStringMap converts a map of stick.Value to a map of string
func makeStringMap(val map[string]stick.Value) (map[string]string, error) {
	target := make(map[string]string)
	for key, val := range val {
		target[key] = stick.CoerceString(val)
	}
	return target, nil
}
