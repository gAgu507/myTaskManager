package cmd

import (
	"fmt"

	"github.com/gAgu507/myTaskManager/internal"
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
				Name:        args[0],
				Description: args[1],
				Done:        status,
			}
			match := CheckForMatch(task.Name)
			fmt.Println("name checked")
			if match {
				fmt.Println("Name of task already exists!")
				return
			} else {
				internal.Add(&task)
			}
		},
	}
)

func CheckForMatch(name string) bool {
	for _, task := range internal.Tasks {
		if task.Name == name {
			return true
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(addCmd)
}
