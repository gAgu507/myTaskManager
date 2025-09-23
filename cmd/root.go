package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gAgu507/myTaskManager/internal"
	"github.com/spf13/cobra"
)

var (
	filename = "internal/storage.json"
	rootCmd  = &cobra.Command{
		Use:   "tm",
		Short: "Basic version of task manager",
		Long:  "This CLI program is my first project on Go. It's a task manager that has basic functions, such as adding, deleting, listing tasks.",
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initApp)
	cobra.OnFinalize(finalizeApp)
}

func initApp() {
	if err := setStruct(); err != nil {
		fmt.Println("Error! Couldn't set the struct while initializing.")
	}
}

func setStruct() error {
	data, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error while reading a file: %v", err)
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &internal.Tasks); err != nil {
			return fmt.Errorf("error while parsing existing file %v", err)
		}
	}
	return nil
}

func finalizeApp() {
	newData, err := json.MarshalIndent(internal.Tasks, "", " ")
	if err != nil {
		fmt.Printf("error marshaling task to json format. %v", err)
	}
	if err := os.WriteFile(filename, newData, 0644); err != nil {
		fmt.Printf("error writing into a json file. %v", err)
	}
}
