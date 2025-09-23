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
	return fmt.Sprintf("%v - %v. | %v", t.Name, t.Description, status)
}

func Check(name string) {
	for _, task := range Tasks {
		fmt.Printf("Got: %v - Have: %v\n", task.Name, name)
		if task.Name == name {
			fmt.Println("Entered equality scope")
			if task.Done {
				fmt.Println("Task is already done!")
				return
			} else {
				fmt.Println("Entered equality scope, where it should be changed")
				task.Done = true
				fmt.Println(task.Done)
			}
		}
	}
}

func Uncheck(name string) {
	for _, task := range Tasks {
		if task.Name == name {
			if !task.Done {
				fmt.Println("Task is not done yet!")
				return
			} else {
				task.Done = false
			}
		}
	}
}
