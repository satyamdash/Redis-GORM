package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
