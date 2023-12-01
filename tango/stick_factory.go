package tango

import (
	"github.com/tyler-sommer/stick"
	"github.com/tyler-sommer/stick/twig/filter"
	"log"
	"os"
)

// CreateStickWithCwd creates a new stick.Env with the default filters and the filters defined in this package
// It uses the current working directory to the filesystem loader
func CreateStickWithCwd() *stick.Env {
	d, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return CreateStickWithWorkDir(d)
}

// CreateStickWithWorkDir creates a new stick.Env with the default filters and the filters defined in this package
func CreateStickWithWorkDir(workDir string) *stick.Env {
	loader := stick.NewFilesystemLoader(workDir)
	env := stick.New(loader)
	env.Filters = mergeMaps(filter.TwigFilters(), CreateFilters())
	return env
}

// mergeMaps merges two maps of string to stick.Filter
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
