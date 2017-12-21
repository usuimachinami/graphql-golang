package model

import "time"

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
