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
	Long:  "Save a bearer token on the current environment. It is sent as Authorization: Bearer ... on requests that use a path like /users.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]

		cfg, err := config.Load()
		if err != nil {
			formatter.Fail(err)
		}

		if cfg.CurrentEnv == "" {
			formatter.Fail(fmt.Errorf("no environment selected"))
		}

		env, exists := cfg.Environments[cfg.CurrentEnv]
		if !exists {
			formatter.Fail(fmt.Errorf("current environment not found"))
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
