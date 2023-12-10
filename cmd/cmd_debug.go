package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// addDebugOption adds the --debug option to the command
func addDebugOption(cmd *cobra.Command, usage string) {
	cmd.Flags().BoolP("debug", "", false, "Debug mode: "+usage)
}

// isInDebugMode returns true if the --debug flag is set
func isInDebugMode(cmd *cobra.Command) bool {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return false
	}
	return debug
}

// debugPrintf prints the format string and arguments if the --debug flag is set
func debugPrintf(cmd *cobra.Command, format string, a ...interface{}) {
	if isInDebugMode(cmd) {
		fmt.Printf(format, a...)
	}
}

// debugPrintln prints the arguments if the --debug flag is set
func debugPrintln(cmd *cobra.Command, a ...interface{}) {
	if isInDebugMode(cmd) {
		fmt.Println(a...)
	}
}
