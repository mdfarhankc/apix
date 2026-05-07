package auth

import (
	"fmt"

	"github.com/mdfarhankc/apix/internal/config"
	"github.com/mdfarhankc/apix/internal/formatter"
	"github.com/spf13/cobra"
)

var bearerCmd = &cobra.Command{
	Use:   "bearer [token]",
	Short: "Set bearer token",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]

		cfg, err := config.Load()
		if err != nil {
			formatter.Fail(err)
		}

		if cfg.CurrentEnv == "" {
			fmt.Println("No environment selected.")
			return
		}

		env, exists := cfg.Environments[cfg.CurrentEnv]
		if !exists {
			fmt.Println("Current environment not found.")
			return
		}

		env.BearerToken = token
		cfg.Environments[cfg.CurrentEnv] = env

		err = config.Save(cfg)
		if err != nil {
			formatter.Fail(err)
		}

		fmt.Println("Bearer token saved.")
	},
}

func init() {
	Command.AddCommand(bearerCmd)
}
