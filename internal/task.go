package internal

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Name        string
	Description string
	IsCompleted bool
}

type TaskStorage struct {
	Tasks []Task
}

func NewStorage() *TaskStorage {
	return &TaskStorage{}
}

func (ts *TaskStorage) AddTask(name, desc string) {
	if name == "" {
		fmt.Println("Task must have a name.")
		return
	}
	if desc == "" {
		desc = "No description"
	}
	ts.Tasks = append(ts.Tasks, Task{
		ID:          uuid.New(),
		Name:        name,
		Description: desc,
		IsCompleted: false,
	})
}

func (ts *TaskStorage) ListTasks() {
	for id, task := range ts.Tasks {
		fmt.Println(id+1, task)
	}
}

func (t Task) String() string {
	var status string
	if t.IsCompleted {
		status = "Done"
	} else {
		status = "To be Done"
	}
	return fmt.Sprintf("%v - %v. | %v", t.Name, t.Description, status)
}
