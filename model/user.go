package model

import (
	"time"
)

type User struct {
	Id        int64      `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Favorite  []Favorite
}