package dtos

type UserVtt struct {
	GenericLog
	Username string `json:"username" validate:"required,min=3,max=255"`
}
