package go_tasks

import "time"

type Task struct {
	Uuid        string
	Title       string
	Description string
	DueDate     time.Time
}

var tasks map[string]*Task

func getTasks() map[string]*Task {
	return tasks
}

func getTask(uuid string) *Task {
	return tasks[uuid]
}

func createTask(uuid string, title string, description string, dueDate time.Time) *Task {
	task := Task{
		Uuid:        uuid,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}
	tasks[task.Uuid] = &task
	return getTask(task.Uuid)
}
