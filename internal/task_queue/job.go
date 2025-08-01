package task_queue

import "fmt"

type Job struct {
	description string
	task        func() error
	status      string
	msg         string
}

func (j *Job) Execute() {
	fmt.Println("Executing job: ", j.description)
	j.status = "running"
	err := j.task()
	if err != nil {
		j.status = "error"
		j.msg = err.Error()
	} else {
		j.status = "done"
	}
	fmt.Printf("Job results: %s\nStatus: %s\nMsg: %s\n", j.description, j.status, j.msg)
}

func NewJob(description string, task func() error) *Job {
	return &Job{
		description: description,
		task:        task,
		status:      "pending",
		msg:         "",
	}
}
