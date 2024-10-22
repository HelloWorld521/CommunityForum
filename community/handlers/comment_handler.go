package handlers

import (
	"community/models"
	"community/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentHandler 评论处理器
type CommentHandler struct {
	CommentService *services.CommentService
}

// CreateComment 处理创建评论请求
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求负载"})
		return
	}
	if err := h.CommentService.CreateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// GetCommentsByPostID 处理获取帖子评论请求
func (h *CommentHandler) GetCommentsByPostID(c *gin.Context) {
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseUint(postIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的帖子ID格式"})
		return
	}
	comments, err := h.CommentService.GetCommentsByPostID(uint(postID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comments)
}

// UpdateComment 处理更新评论请求
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	commentIDStr := c.Param("id")
	var comment models.Comment

	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID格式"})
		return
	}

	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求负载"})
		return
	}
	comment.ID = uint(commentID) // 确保更新正确的评论
	if err := h.CommentService.UpdateComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, comment)
}

// DeleteComment 处理删除评论请求
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	commentIDStr := c.Param("id")
	commentID, err := strconv.ParseUint(commentIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的评论ID格式"})
		return
	}

	if err := h.CommentService.DeleteComment(uint(commentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
