package stick_filters

import (
	"github.com/tyler-sommer/stick"
	"log"
)

// fatalityFilter logs a message and exits the program
func fatalityFilter(
	ctx stick.Context,
	val stick.Value,
	args ...stick.Value,
) stick.Value {
	message := stick.CoerceString(val)
	log.Fatal(message)
	return nil
}
