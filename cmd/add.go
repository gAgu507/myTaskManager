package cmd

import (
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add name [description]",
		Short: "Adds a new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskStorage.AddTask(args[0], args[1])
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
