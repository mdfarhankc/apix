package env

import (
	"fmt"

	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [name]",
	Short: "Switch to an environment",
	Long:  "Make this environment the active one. After 'env use staging', a request to /users hits the staging base URL.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		cfg, err := config.Load()
		if err != nil {
			formatter.Fail(err)
		}

		_, exists := cfg.Environments[name]
		if !exists {
			fmt.Printf(
				"Environment '%s' not found.\n",
				name,
			)
			return
		}

		cfg.CurrentEnv = name

		err = config.Save(cfg)
		if err != nil {
			formatter.Fail(err)
		}

		fmt.Printf(
			"Current environment set to '%s'.\n",
			name,
		)
	},
}

func init() {
	Command.AddCommand(useCmd)
}
