package user_vtt

import (
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func (r *Repo) Create(user UserVttModel) error {
	return r.db.Create(&user).Error
}

func NewUserVttRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}
