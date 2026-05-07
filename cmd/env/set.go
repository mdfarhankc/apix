package env

import (
	"fmt"

	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set [name] [base_url]",
	Short: "Create or update an environment",
	Long:  "Save a base URL under a name. Re-running with the same name overwrites it.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		baseURL := args[1]

		cfg, err := config.Load()
		if err != nil {
			formatter.Fail(err)
		}

		cfg.Environments[name] = config.Environment{
			BaseURL: baseURL,
		}

		err = config.Save(cfg)
		if err != nil {
			formatter.Fail(err)
		}

		fmt.Printf(
			"Environment '%s' saved.\n",
			name,
		)
	},
}

func init() {
	Command.AddCommand(setCmd)
}
