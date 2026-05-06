package cmd

import (
	"fmt"

	"github.com/mdfarhankc/apix/internal/client"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [url]",
	Short: "Send a GET request",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]
		fmt.Println("GET:", url)

		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("%d (%v)\n\n", resp.StatusCode, resp.Duration)
		fmt.Println(string(resp.Body))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
