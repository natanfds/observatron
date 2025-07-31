package dtos

import "github.com/natanfds/observatron/types"

type GenericLog struct {
	Level     types.LogLevel `json:"level" validate:"required,min=4,max=4"`
	Action    string         `json:"action" validate:"required,min=4,max=255"`
	Message   string         `json:"message" validate:"required,min=10,max=255"`
	CreatedAt string         `json:"created_at" validate:"required"`
}
