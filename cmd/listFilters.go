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

// listFiltersCmd represents the listFilters command
var listFiltersCmd = &cobra.Command{
	Use:   "stick:list-filters",
	Short: "Lists filters available to templates",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		stickEnv := tango.CreateStick()
		// Extract keys into a slice and sort
		keys := make([]string, 0, len(stickEnv.Filters))
		for key := range stickEnv.Filters {
			keys = append(keys, key)
		}
		sort.Strings(keys)

		fmt.Println("Filters:")
		for _, f := range keys {
			fmt.Println(f)
		}

	},
}

func init() {
	rootCmd.AddCommand(listFiltersCmd)
}
