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
}

type Title struct {
	PublisherId string     `json:"publisher_id"`
	TitleId     string     `json:"title_id"`
	FullTitleId string     `gorm:"primary_key" json:"id"`
	TitleName   string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Story       []Story
}

type Story struct {
	FullTitleId string     `json:"full_title_id"`
	StoryId     string     `json:"story_id"`
	FullStoryId string     `gorm:"primary_key" json:"id"`
	StoryName   string     `json:"name"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Title       Title      `gorm:"ForeignKey:FullTitleId"`
}
