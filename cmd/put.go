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
		queryFlags, _ := cmd.Flags().GetStringArray("query")
		runner.Run(runner.Options{
			Method:      "PUT",
			RawURL:      args[0],
			Body:        []byte(body),
			RawHeaders:  headerFlags,
			RawQuery:    queryFlags,
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
	putCmd.Flags().StringArrayP(
		"query",
		"q",
		[]string{},
		"Query parameter (key=val)",
	)

	rootCmd.AddCommand(putCmd)
}
