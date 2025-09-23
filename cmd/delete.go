package cmd

import (
	"github.com/gAgu507/myTaskManager/internal"
	"github.com/spf13/cobra"
)

var (
	deleteCmd = &cobra.Command{
		Use:   "delete name",
		Short: "Deletes a task by name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			internal.Delete(name)
		},
	}
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}
