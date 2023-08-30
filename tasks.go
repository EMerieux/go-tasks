package go_tasks

import "time"

type Task struct {
	Uuid        string
	Title       string
	Description string
	DueDate     time.Time
}

var tasks map[string]*Task

func GetTasks() map[string]*Task {
	return tasks
}

func GetTask(uuid string) *Task {
	return tasks[uuid]
}

func CreateTask(uuid string, title string, description string, dueDate time.Time) *Task {
	task := Task{
		Uuid:        uuid,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}
	tasks[task.Uuid] = &task
	return tasks[task.Uuid]
}
