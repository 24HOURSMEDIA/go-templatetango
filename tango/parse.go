package tango

import (
	"bytes"
	"github.com/tyler-sommer/stick"
)

func Parse(template string, params *map[string]stick.Value) (string, error) {
	stickEnv := CreateStick()
	buf := new(bytes.Buffer)
	if err := stickEnv.Execute(
		template,
		buf,
		*params,
	); err != nil {
		return "", err
	}
	return buf.String(), nil
}
