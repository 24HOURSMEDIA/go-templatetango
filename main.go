/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"templatetango/cmd"
	"templatetango/tango"
)

func main() {
	err := tango.LoadDotEnv(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	cmd.Execute()
}
