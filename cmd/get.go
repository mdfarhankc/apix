package cmd

import (
	"github.com/mdfarhankc/apix/internal/runner"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Send a GET request",
	Long:  "Send a GET request to a URL. Add headers with -H, repeat the flag for more.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		queryFlags, _ := cmd.Flags().GetStringArray("query")
		runner.Run(runner.Options{
			Method:     "GET",
			RawURL:     args[0],
			RawHeaders: headerFlags,
			RawQuery:   queryFlags,
		})
	},
}

func init() {
	getCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)
	getCmd.Flags().StringArrayP(
		"query",
		"q",
		[]string{},
		"Query parameter (key=val)",
	)
	rootCmd.AddCommand(getCmd)
}
