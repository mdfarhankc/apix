package cmd

import (
	"github.com/mdfarhankc/apix/internal/runner"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [url]",
	Short: "Send a DELETE request",
	Long:  "Send a DELETE request to a URL. Add headers with -H, repeat the flag for more.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		queryFlags, _ := cmd.Flags().GetStringArray("query")
		runner.Run(runner.Options{
			Method:     "DELETE",
			RawURL:     args[0],
			RawHeaders: headerFlags,
			RawQuery:   queryFlags,
		})
	},
}

func init() {
	deleteCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)
	deleteCmd.Flags().StringArrayP(
		"query",
		"q",
		[]string{},
		"Query parameter (key=val)",
	)
	rootCmd.AddCommand(deleteCmd)
}
