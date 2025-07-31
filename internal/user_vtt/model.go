package user_vtt

import (
	"github.com/natanfds/observatron/internal/generics"
)

type UserVttModel struct {
	generics.GenericLog
	Username string `gorm:"index;not null"`
}
