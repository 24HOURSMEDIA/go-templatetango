package tango

import (
	"github.com/tyler-sommer/stick"
	"os"
	"strings"
)

// CreateParams creates a map of environment variables
func CreateParams() *map[string]stick.Value {
	// get all environment variables as a map of strings
	vars := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		if len(pair) == 2 {
			vars[pair[0]] = pair[1]
		}
	}
	// create a new map like map[string]stick.Value{"name": "Tyler"}
	params := make(map[string]stick.Value)
	for key, val := range vars {
		params[key] = stick.Value(val)
	}

	return &params
}
