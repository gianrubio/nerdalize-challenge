package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:   "task",
	Short: "Create and list a task",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
