package handlers

import (
	"community/models"
	"community/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// PostHandler 帖子处理器
type PostHandler struct {
	PostService *services.PostService
}

// CreatePost 处理创建帖子请求
func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := h.PostService.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, post)
}

// GetPostByID 处理获取帖子请求
func (h *PostHandler) GetPostByID(c *gin.Context) {
	idStr := c.Param("id")

	// 转换字符串为 uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	post, err := h.PostService.GetPostByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "帖子未找到"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// GetAllPosts 处理获取所有帖子请求
func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.PostService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// UpdatePost 处理更新帖子请求
func (h *PostHandler) UpdatePost(c *gin.Context) {
	idStr := c.Param("id")
	var post models.Post

	// 转换字符串为 uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	post.ID = uint(id) // 确保更新正确的帖子
	if err := h.PostService.UpdatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeletePost 处理删除帖子请求
func (h *PostHandler) DeletePost(c *gin.Context) {
	idStr := c.Param("id")

	// 转换字符串为 uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的ID格式"})
		return
	}

	if err := h.PostService.DeletePost(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
