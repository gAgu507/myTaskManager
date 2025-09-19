package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	filename = "internal/storage.json"
	rootCmd  = &cobra.Command{
		Use:   "tm",
		Short: "Basic version of task manager",
		Long:  "This CLI program is my first project on Go. It's a task manager that has basic functions, such as adding, deleting, listing task.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
