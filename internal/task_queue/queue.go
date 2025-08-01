package task_queue

import (
	"errors"
	"fmt"
)

type TaskQueue struct {
	jobQueue     chan Job
	closeQueue   chan bool
	queueMaxSize int
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		jobQueue:   make(chan Job, 1),
		closeQueue: make(chan bool),
	}
}

func (t *TaskQueue) Start() {
	go func() {
		fmt.Println("Starting job queues")
		for {
			select {
			case job := <-t.jobQueue:
				job.Execute()
			case <-t.closeQueue:
				if t.QueueLength() == 0 {
					fmt.Println("Closing job queue")
					return
				}
			}
		}
	}()
}

func (t *TaskQueue) Stop() {
	t.closeQueue <- true
}

// Função que adiciona task a fila
func (t *TaskQueue) AddJob(job Job) error {
	select {
	case <-t.closeQueue:
		return errors.New("queue closed")
	default:
		t.jobQueue <- job
		return nil
	}
}

func (t *TaskQueue) QueueLength() int {
	return len(t.jobQueue)
}

func (t *TaskQueue) IsQueueClosed() bool {
	return <-t.closeQueue
}
