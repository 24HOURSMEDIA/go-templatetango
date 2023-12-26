package stick_filters

import (
	"fmt"
	"github.com/tyler-sommer/stick"
	"log"
	"strconv"
)

// extractObjectsFilter extracts objects from variables in the current scope
// by applying a mapping to them. Only if the scope contains the first key,
// the mapping is applied and added to the result.
func extractObjectsFilter(
	ctx stick.Context,
	val stick.Value,
	args ...stick.Value,
) stick.Value {
	mapping, err := makeStringMap(val.(map[string]stick.Value))
	if err != nil || len(mapping) == 0 {
		return nil
	}
	if len(args) < 2 {
		log.Fatal("extractObjectsFilter needs at least 2 arguments: count, requiredKey, defaultVal")
	}
	count := int(stick.CoerceNumber(args[0]))
	requiredKey := stick.CoerceString(args[1])

	defaultVal := stick.Value(nil)
	if len(args) > 2 {
		defaultVal = args[2]
	}

	var suffices []string
	suffices = append(suffices, "")

	for i := 0; i < count; i++ {
		suffices = append(suffices, strconv.Itoa(i), "_"+strconv.Itoa(i))
	}
	var result []stick.Value
	scopeVars := ctx.Scope().All()
	for _, suffix := range suffices {
		requiredKeyWithSuffix := requiredKey + suffix
		fmt.Println("SUFFIX", requiredKeyWithSuffix)
		if _, ok := scopeVars[requiredKeyWithSuffix]; ok {
			result = append(result, applyMappingWithSuffix(mapping, scopeVars, suffix, defaultVal))
		}
	}
	return stickify(result)
}
