/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"templatetango/tango"
)

// parseFileCmd represents the parseFile command
var parseFileCmd = &cobra.Command{
	Use:   "parse:file",
	Short: "Parses a twig-like file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("parse:template requires exactly one argument")
			return
		}
		filePath := args[0]
		// load the template or return an error
		template, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
			return
		}

		parsed, err := tango.Parse(string(template), tango.CreateParams())
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = fmt.Fprintf(os.Stdout, parsed)
		if err != nil {
			log.Fatal(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(parseFileCmd)
}
