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
		runner.Run(runner.Options{
			Method:     "DELETE",
			RawURL:     args[0],
			RawHeaders: headerFlags,
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
	rootCmd.AddCommand(deleteCmd)
}
