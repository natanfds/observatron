package user_vtt

import (
	"time"

	"github.com/natanfds/observatron/dtos"
	"github.com/natanfds/observatron/internal/generics"
	"github.com/natanfds/observatron/internal/task_queue"
)

type Service struct {
	repo      *Repo
	taskQueue *task_queue.TaskQueue
}

func (s *Service) Create(data dtos.UserVtt) error {
	timeStamp, err := time.Parse(time.RFC3339, data.CreatedAt)
	if err != nil {
		return err
	}

	job := task_queue.NewJob("Salvando Log de Usuário VTT", func() error {
		return s.repo.Create(UserVttModel{
			GenericLog: generics.GenericLog{
				Level:        data.Level,
				Message:      data.Message,
				LogCreatedAt: timeStamp,
				Action:       data.Action,
			},
			Username: data.Username,
		})
	})
	err = s.taskQueue.AddJob(*job)
	if err != nil {
		return err
	}

	return nil
}

func NewUserVttService(repo *Repo, taskQueue *task_queue.TaskQueue) *Service {
	return &Service{repo: repo, taskQueue: taskQueue}
}
