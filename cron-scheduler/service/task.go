package service

import (
	"fmt"
)

type Task struct {
	TaskID int
	Type   string
}

func (t *Task) Execute() {
	fmt.Println(fmt.Sprintf("Task %d running", t.TaskID))
}
