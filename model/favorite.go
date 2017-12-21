package model

import "time"

type Favorite struct {
	Id          int        `gorm:"primary_key" json:"id"`
	FullTitleId string     `json:"full_title_id"`
	UserId      int        `json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	User        User       `gorm:"ForeignKey:UserId"`
}