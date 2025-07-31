package user_vtt

import (
	"time"

	"github.com/natanfds/observatron/dtos"
	"github.com/natanfds/observatron/internal/generics"
)

type Service struct {
	repo *Repo
}

func (s Service) Create(data dtos.UserVtt) error {
	timeStamp, err := time.Parse(time.RFC3339, data.CreatedAt)
	if err != nil {
		return err
	}

	return s.repo.Create(UserVttModel{

		GenericLog: generics.GenericLog{
			Level:        data.Level,
			Message:      data.Message,
			LogCreatedAt: timeStamp,
			Action:       data.Action,
		},
		Username: data.Username,
	})
}

func NewUserVttService(repo *Repo) *Service {
	return &Service{repo: repo}
}
