package cmd

import (
	"github.com/grahamplata/sixers/handlers"
	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Fetches the next available sixers game date and time.",
	Long:  "Fetches the next available sixers game date and time.",
	Run: func(cmd *cobra.Command, args []string) {
		handlers.Next()
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
