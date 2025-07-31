package generics

import (
	"time"

	"gorm.io/gorm"

	"github.com/natanfds/observatron/types"
)

type GenericLog struct {
	gorm.Model
	Level        types.LogLevel `gorm:"index;not null"`
	Message      string         `gorm:"not null"`
	LogCreatedAt time.Time      `gorm:"index;not null"`
	Action       string         `gorm:"index;not null"`
}
