package internal

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}

func Add(filename string, t *Task) error {
	var tasks []Task
	data, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error while reading a file: %v", err)
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return fmt.Errorf("error while parsing existing file %v", err)
		}
	}

	tasks = append(tasks, *t)

	newData, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return fmt.Errorf("error marshaling task to json format. %v", err)
	}
	if err := os.WriteFile(filename, newData, 0644); err != nil {
		return fmt.Errorf("error writing into a json file. %v", err)
	}

	fmt.Println("Task's added.")
	return nil
}

func List(filename string) error {
	var tasks []Task
	data, err := os.ReadFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error while reading a file: %v", err)
	}

	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return fmt.Errorf("error while parsing existing file %v", err)
		}
	}
	for i, task := range tasks {
		fmt.Println(i+1, task)
	}
	return nil
}

func (t Task) String() string {
	var status string
	if t.Done {
		status = "Done"
	} else {
		status = "To Do"
	}
	return fmt.Sprintf("%v - %v. | %v", t.Name, t.Description, status)
}
