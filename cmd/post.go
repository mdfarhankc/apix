package cmd

import (
	"github.com/mdfarhankc/apix/internal/runner"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Send a POST request",
	Long:  "Send a POST request with a body (-d) and optional headers (-H). Defaults to JSON.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		body, _ := cmd.Flags().GetString("data")
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		queryFlags, _ := cmd.Flags().GetStringArray("query")
		runner.Run(runner.Options{
			Method:      "POST",
			RawURL:      args[0],
			Body:        []byte(body),
			RawHeaders:  headerFlags,
			RawQuery:    queryFlags,
			ContentType: "application/json",
		})
	},
}

func init() {
	postCmd.Flags().StringP(
		"data",
		"d",
		"",
		"Request body",
	)
	postCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)
	postCmd.Flags().StringArrayP(
		"query",
		"q",
		[]string{},
		"Query parameter (key=val)",
	)

	rootCmd.AddCommand(postCmd)
}
