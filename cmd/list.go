package cmd

import (
	"github.com/gAgu507/myTaskManager/internal"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all tasks",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			internal.List(filename)
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
