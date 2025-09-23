package cmd

import (
	"github.com/gAgu507/myTaskManager/internal"
	"github.com/spf13/cobra"
)

var (
	checkCmd = &cobra.Command{
		Use:   "check",
		Short: "checks the task",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			internal.Check(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(checkCmd)
}
