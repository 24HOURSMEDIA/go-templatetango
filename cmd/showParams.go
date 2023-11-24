/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"templatetango/tango"
)

// showParamsCmd represents the showParams command
var showParamsCmd = &cobra.Command{
	Use:   "stick:show-params",
	Short: "Show parameter names passed to templates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Parameters:")
		params := *tango.CreateParams()
		// Extract keys into a slice and sort
		keys := make([]string, 0, len(params))
		for key := range params {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			fmt.Printf("%s\n", key)
		}
	},
}

func init() {
	rootCmd.AddCommand(showParamsCmd)
}
