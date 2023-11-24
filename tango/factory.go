package tango

import (
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig/filter"
)

func CreateStick() *stick.Env {
	env := stick.New(nil)
	env.Filters = mergeMaps(filter.TwigFilters(), CreateFilters())
	return env
}

func mergeMaps(map1, map2 map[string]stick.Filter) map[string]stick.Filter {
	merged := make(map[string]stick.Filter)
	for key, value := range map1 {
		merged[key] = value
	}
	// Add all key-value pairs from map2 to the merged map, overwriting any existing keys
	for key, value := range map2 {
		merged[key] = value
	}
	return merged
}
