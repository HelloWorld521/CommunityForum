package models

import (
	"gorm.io/gorm"
)

// Comment 评论模型
type Comment struct {
	gorm.Model
	PostID  uint   `json:"post_id" gorm:"not null"`
	UserID  uint   `json:"user_id" gorm:"not null"`
	Content string `json:"content" gorm:"not null"`
}
