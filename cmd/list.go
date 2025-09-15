package cmd

import (
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all tasks",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			taskStorage.ListTasks()
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
