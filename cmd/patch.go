package cmd

import (
	"github.com/mdfarhankc/apix/internal/runner"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch [url]",
	Short: "Send a PATCH request",
	Long:  "Send a PATCH request with a body (-d) and optional headers (-H). Defaults to JSON.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		body, _ := cmd.Flags().GetString("data")
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		runner.Run(runner.Options{
			Method:      "PATCH",
			RawURL:      args[0],
			Body:        []byte(body),
			RawHeaders:  headerFlags,
			ContentType: "application/json",
		})
	},
}

func init() {
	patchCmd.Flags().StringP(
		"data",
		"d",
		"",
		"Request body",
	)
	patchCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)

	rootCmd.AddCommand(patchCmd)
}
