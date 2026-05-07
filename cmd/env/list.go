package env

import (
	"fmt"

	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved environments",
	Long:  "Show every environment and its base URL. The current environment is the one used when you call a path like /users.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			formatter.Fail(err)
		}

		if len(cfg.Environments) == 0 {
			fmt.Println("No environments found.")
			return
		}

		for name, env := range cfg.Environments {
			fmt.Printf(
				"%s\t%s\n",
				name,
				env.BaseURL,
			)
		}
	},
}

func init() {
	Command.AddCommand(listCmd)
}
