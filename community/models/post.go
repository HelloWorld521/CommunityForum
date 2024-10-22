package models

import (
	"gorm.io/gorm"
)

// Post 帖子模型
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"user_id"`
}
