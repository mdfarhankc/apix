package env

import "github.com/spf13/cobra"

var Command = &cobra.Command{
	Use:   "env",
	Short: "Manage environments",
	Long:  "Save base URLs as named environments and switch between them, so you can call /users instead of typing the full host every time.",
}
