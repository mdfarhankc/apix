package cmd

import (
	"github.com/mdfarhankc/apix/internal/runner"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put [url]",
	Short: "Send a PUT request",
	Long:  "Send a PUT request with a body (-d) and optional headers (-H). Defaults to JSON.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		body, _ := cmd.Flags().GetString("data")
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		runner.Run(runner.Options{
			Method:      "PUT",
			RawURL:      args[0],
			Body:        []byte(body),
			RawHeaders:  headerFlags,
			ContentType: "application/json",
		})
	},
}

func init() {
	putCmd.Flags().StringP(
		"data",
		"d",
		"",
		"Request body",
	)
	putCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)

	rootCmd.AddCommand(putCmd)
}
