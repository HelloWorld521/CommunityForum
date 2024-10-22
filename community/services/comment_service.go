package services

import (
	"community/models"
	"gorm.io/gorm"
)

// CommentService 评论服务
type CommentService struct {
	DB *gorm.DB
}

// CreateComment 创建评论
func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.DB.Create(comment).Error
}

// GetCommentsByPostID 根据帖子ID获取评论
func (s *CommentService) GetCommentsByPostID(postID uint) ([]models.Comment, error) {
	var comments []models.Comment
	err := s.DB.Where("post_id = ?", postID).Find(&comments).Error
	return comments, err
}

// UpdateComment 更新评论
func (s *CommentService) UpdateComment(comment *models.Comment) error {
	return s.DB.Save(comment).Error
}

// DeleteComment 删除评论
func (s *CommentService) DeleteComment(commentID uint) error {
	return s.DB.Delete(&models.Comment{}, commentID).Error
}
