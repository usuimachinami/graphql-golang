package model

import "time"

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