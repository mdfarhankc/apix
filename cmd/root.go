package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "apix",
	Short:   "A simple HTTP client for the terminal",
	Long:    "APIX sends HTTP requests and prints the response — colorized, timed, and pretty when it's JSON.",
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
