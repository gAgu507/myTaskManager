package internal

import (
	"fmt"
)

var (
	Tasks []Task
)

type Task struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func Add(t *Task) {
	Tasks = append(Tasks, *t)
}

func Delete(name string) {
	for i, task := range Tasks {
		if task.Name == name {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
}

func List() {
	for i, task := range Tasks {
		fmt.Println(i+1, task)
	}
}

func (t Task) String() string {
	var status string
	if t.Done {
		status = "Done"
	} else {
		status = "To Do"
	}
	return fmt.Sprintf("%v - %v. | %v\n", t.Name, t.Description, status)
}

func Check(name string) {
	for i, task := range Tasks {
		if task.Name == name {
			if task.Done {
				fmt.Println("Task is already done!")
				return
			} else {
				Tasks[i].Done = true
			}
		}
	}
}

func Uncheck(name string) {
	for i, task := range Tasks {
		if task.Name == name {
			if !task.Done {
				fmt.Println("Task is not done yet!")
				return
			} else {
				Tasks[i].Done = false
			}
		}
	}
}
