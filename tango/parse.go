package tango

import (
	"bytes"
	"github.com/tyler-sommer/stick"
)

// Parse parses a twig template with the given parameters
func Parse(template string, params *map[string]stick.Value) (string, error) {
	stickEnv := CreateStickWithCwd()
	return ParseWithStickEnv(template, params, stickEnv)
}

// ParseWithStickEnv parses a twig template with the given parameters and stick.Env
func ParseWithStickEnv(template string, params *map[string]stick.Value, stickEnv *stick.Env) (string, error) {
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
