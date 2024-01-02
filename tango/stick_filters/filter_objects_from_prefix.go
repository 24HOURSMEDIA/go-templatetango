package stick_filters

import (
	"github.com/tyler-sommer/stick"
	"strconv"
	"strings"
)

// cloneMap returns a copy of a map
func cloneMap(original map[string]interface{}) map[string]interface{} {
	cloned := make(map[string]interface{})
	for key, value := range original {
		cloned[key] = value
	}
	return cloned
}

// objectsFromPrefix returns a slice of objects from a map of vars
// that have a prefix and a suffix. The prefix is the same for all
// objects, the suffix is a number from 0 to maxSuffix.
func objectsFromPrefix(vars map[string]interface{}, prefix string, maxSuffix int, defaultObj map[string]interface{}) []interface{} {
	sep := "_"
	mappedResults := map[string]map[string]interface{}{}
	var prefixes []string
	prefixes = append(prefixes, prefix+sep)
	for i := 0; i <= maxSuffix; i++ {
		prefixes = append(prefixes, prefix+strconv.Itoa(i)+sep)
	}
	for key, val := range vars {
		if !strings.HasPrefix(key, prefix) {
			continue
		}
		for _, prefixToDetect := range prefixes {
			if strings.HasPrefix(key, prefixToDetect) {
				if _, ok := mappedResults[prefixToDetect]; !ok {
					mappedResults[prefixToDetect] = cloneMap(defaultObj)
				}
				truncatedKey := strings.TrimPrefix(key, prefixToDetect)
				mappedResults[prefixToDetect][truncatedKey] = val
			}
		}
	}

	// add by order of prefixes
	var results []interface{}
	for _, prefixToDetect := range prefixes {
		if _, ok := mappedResults[prefixToDetect]; ok {
			results = append(results, mappedResults[prefixToDetect])
		}
	}
	return results
}

func ObjectsFromPrefixFilter(
	ctx stick.Context,
	val stick.Value,
	args ...stick.Value,
) stick.Value {
	if len(args) < 1 {
		return nil
	}
	vars := ctx.Scope().All()
	prefix := stick.CoerceString(val)
	maxSuffix := int(stick.CoerceNumber(args[0]))
	defaultObj := map[string]stick.Value{}
	if len(args) > 1 {
		defaultObj = args[1].(map[string]stick.Value)
	}
	return stickify(objectsFromPrefix(destickifyMap(vars), prefix, maxSuffix, destickifyMap(defaultObj)))
}
