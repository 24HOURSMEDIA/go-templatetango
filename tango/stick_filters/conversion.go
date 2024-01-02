package stick_filters

import (
	"errors"
	"github.com/tyler-sommer/stick"
	"reflect"
	"strconv"
	"strings"
)

// stickify converts all elements of a slice or map to stick.Value
// recursively.
func stickify(v interface{}) stick.Value {
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		// Create a new slice where each element is a stick.Value
		length := rv.Len()
		slice := make([]stick.Value, length)
		for i := 0; i < length; i++ {
			slice[i] = stickify(rv.Index(i).Interface())
		}
		return slice
	case reflect.Map:
		// Create a new map where each value is a stick.Value
		mapType := reflect.MapOf(rv.Type().Key(), reflect.TypeOf((*stick.Value)(nil)).Elem())
		newMap := reflect.MakeMap(mapType)
		for _, key := range rv.MapKeys() {
			newMap.SetMapIndex(key, reflect.ValueOf(stickify(rv.MapIndex(key).Interface())))
		}
		return newMap.Interface()
	default:
		// It's not a slice or a map, so just return it as is.
		return v
	}
}

func destickifyMap(v map[string]stick.Value) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range v {
		result[key] = value
	}
	return result
}

// boolify converts a value to a bool. It supports the following types:
// bool, string, int, float32, float64
// Float and int values that do not evaluate to 0 are converted to true.
// String values that can be converted to a float or int are converted to a bool using the above rules.
// It also supports the following strings:
// "true", "yes", "y", "on"
// "false", "no", "n", "off", ""
// It returns an error if the value cannot be converted to a bool.
func boolify(v interface{}) (bool, error) {
	if v == nil {
		return false, nil
	}
	reflectValue := reflect.ValueOf(v)
	returnValue := false
	switch reflectValue.Kind() {
	case reflect.Bool:
		return reflectValue.Bool(), nil
	case reflect.String:
		floatVal, floatErr := strconv.ParseFloat(reflectValue.String(), 64)
		if floatErr == nil {
			return boolify(floatVal)
		}
		intVal, intErr := strconv.ParseInt(reflectValue.String(), 10, 64)
		if intErr == nil {
			return boolify(intVal)
		}
		returnValue = isLowercaseStringInArray(reflectValue.String(), []string{
			"true",
			"yes",
			"y",
			"on",
			"enabled",
			"enable",
		})
		if !returnValue && !isLowercaseStringInArray(reflectValue.String(), []string{
			"false",
			"no",
			"n",
			"off",
			"disabled",
			"disable",
			"",
		}) {
			return false, errors.New("boolify() could not convert string " + reflectValue.String() + " to bool")
		}
		return returnValue, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflectValue.Int() != 0, nil
	case reflect.Float32, reflect.Float64:
		return reflectValue.Float() != 0, nil
	default:
		return false, errors.New("boolify() does not support type " + reflectValue.Kind().String())
	}
}

// isLowercaseStringInArray checks if the lowercase version of str is in the array arr.
func isLowercaseStringInArray(str string, arr []string) bool {
	lowerStr := strings.ToLower(str)
	for _, value := range arr {
		if lowerStr == value {
			return true
		}
	}
	return false
}
