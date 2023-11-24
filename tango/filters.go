package tango

import (
	"github.com/tyler-sommer/stick"
)

func CreateFilters() map[string]stick.Filter {
	return map[string]stick.Filter{
		"say_hello": sayHello,
	}
}

func sayHello(ctx stick.Context, val stick.Value, args ...stick.Value) stick.Value {
	return "Hello " + stick.CoerceString(val) + " from a custom filter!"
}
