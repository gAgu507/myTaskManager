package cmd

import (
	"fmt"
	"os"

	"github.com/gAgu507/myTaskManager/internal"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	taskStorage = internal.NewStorage()
	rootCmd     = &cobra.Command{
		Use:   "tm",
		Short: "Basic version of task manager",
		Long:  "This CLI program is my first project on Go. It's a task manager that has basic functions, such as adding, deleting, listing task.",

		// Run: func(cmd *cobra.Command, args []string) {
		// 	fmt.Println("You started task manager!")
		// },
	}
)

func Execute() {
	// delete
	taskStorage.Tasks = append(taskStorage.Tasks, internal.Task{
		ID:          uuid.New(),
		Name:        "Some name",
		Description: "This desc",
		IsCompleted: true,
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
