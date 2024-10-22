package services

import (
	"community/models"
	"gorm.io/gorm"
)

// PostService 帖子服务
type PostService struct {
	DB *gorm.DB
}

// CreatePost 创建新帖子
func (s *PostService) CreatePost(post *models.Post) error {
	return s.DB.Create(post).Error
}

// GetPosts 获取所有帖子
func (s *PostService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	err := s.DB.Find(&posts).Error
	return posts, err
}

// GetPostByID 根据ID获取帖子
func (s *PostService) GetPostByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := s.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

// UpdatePost 更新帖子
func (s *PostService) UpdatePost(post *models.Post) error {
	return s.DB.Save(post).Error
}

// DeletePost 删除帖子
func (s *PostService) DeletePost(id uint) error {
	return s.DB.Delete(&models.Post{}, id).Error
}
