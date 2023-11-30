package tango

import (
	"github.com/tyler-sommer/stick"
	"reflect"
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
