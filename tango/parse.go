package tango

import (
	"bytes"
	"github.com/tyler-sommer/stick"
)

// Parse parses a twig template with the given parameters
// deprecated: use ParseFileWithWorkDir instead!
func Parse(templateFile string, params *map[string]stick.Value) (string, error) {
	stickEnv := CreateStickWithCwd()
	return ParseWithStickEnv(templateFile, params, stickEnv)
}

// ParseFileWithWorkDir parses a twig template with the given parameters and the given working directory
func ParseFileWithWorkDir(templateFile string, params *map[string]stick.Value, workDir string) (string, error) {
	stickEnv := CreateStickWithWorkDir(workDir)
	return ParseWithStickEnv(templateFile, params, stickEnv)
}

// ParseString parses a template string directly. This does not support 'include' etc.
func ParseString(template string, params *map[string]stick.Value) (string, error) {
	stickEnv := CreateStickStringParser()
	return ParseWithStickEnv(template, params, stickEnv)
}

// ParseWithStickEnv parses a twig template with the given parameters and stick.Env
// If the loader is a filesystem loader, it will use the current working directory and the template is a file.
// If the loader is a string loader, it will use the template as a string.
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
