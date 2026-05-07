package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/mdfarhankc/apix/internal/client"
	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Send a GET request",
	Long:  "Send a GET request to a URL. Add headers with -H, repeat the flag for more.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url, err := config.ResolveURL(args[0])
		if err != nil {
			formatter.Fail(err)
		}

		headerFlags, _ := cmd.Flags().GetStringArray("header")
		headers := client.ParseHeaders(headerFlags)

		fmt.Printf(
			"%s %s\n\n",
			color.CyanString("GET"),
			url,
		)

		request := client.Request{
			Method:  "GET",
			URL:     url,
			Headers: headers,
		}

		resp, err := client.Do(request)
		if err != nil {
			formatter.Fail(err)
		}

		formatter.PrintResponse(resp)
	},
}

func init() {
	getCmd.Flags().StringArrayP(
		"header",
		"H",
		[]string{},
		"Request headers",
	)
	rootCmd.AddCommand(getCmd)
}
