package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mdfarhankc/apix/internal/client"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post [url]",
	Short: "Send a POST request",
	Long:  "Send a POST request with a body (-d) and optional headers (-H). Defaults to JSON.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		body, _ := cmd.Flags().GetString("data")
		headerFlags, _ := cmd.Flags().GetStringArray("header")
		headers := client.ParseHeaders(headerFlags)
		if _, ok := headers["Content-Type"]; !ok {
			headers["Content-Type"] = "application/json"
		}

		fmt.Printf(
			"%s %s\n\n",
			color.CyanString("POST"),
			url,
		)

		request := client.Request{
			Method:  "POST",
			URL:     url,
			Body:    []byte(body),
			Headers: headers,
		}

		resp, err := client.Do(request)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		formatter.PrintResponse(resp)
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

	rootCmd.AddCommand(postCmd)
}
