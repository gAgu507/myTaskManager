package cmd

import (
	"github.com/gAgu507/myTaskManager/internal"
	"github.com/spf13/cobra"
)

var (
	uncheckCmd = &cobra.Command{
		Use:   "uncheck",
		Short: "unchecks the task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.Uncheck(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(uncheckCmd)
}
