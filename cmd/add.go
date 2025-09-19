package cmd

import (
	"fmt"

	"github.com/gAgu507/myTaskManager/internal"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
		Use:   "add name description status",
		Short: "Adds a new task",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			status := false
			if args[2] == "Done" {
				status = true
			}
			task := internal.Task{
				ID:          uuid.New(),
				Name:        args[0],
				Description: args[1],
				Done:        status,
			}

			err := internal.Add(filename, &task)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
}
