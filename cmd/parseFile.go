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
		filePathArg := args[0]
		// ! No info on stdout!
		resolved, err := resolveTemplateFileDirAndPath(cmd, filePathArg)
		if err != nil {
			log.Fatal(err)
			return
		}

		stickEnv := tango.CreateStickWithWorkDir(resolved.AbsDir)
		params := tango.CreateParams()
		parsed, err := tango.ParseWithStickEnv(resolved.Relative, params, stickEnv)
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
	addTemplatesDirOption(parseFileCmd)
	rootCmd.AddCommand(parseFileCmd)
}
