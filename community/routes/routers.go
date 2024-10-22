package routes

import (
	"community/handlers"
	"community/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 初始化服务
	authService := &services.AuthService{DB: db}
	postService := &services.PostService{DB: db}
	commentService := &services.CommentService{DB: db}
	authHandler := &handlers.AuthHandler{AuthService: authService}
	postHandler := &handlers.PostHandler{PostService: postService}
	commentHandler := &handlers.CommentHandler{CommentService: commentService}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)

		// 帖子相关路由
		v1.POST("/posts", postHandler.CreatePost)
		v1.GET("/posts", postHandler.GetAllPosts)
		v1.GET("/posts/:id", postHandler.GetPostByID)
		v1.PUT("/posts/:id", postHandler.UpdatePost)
		v1.DELETE("/posts/:id", postHandler.DeletePost)

		// 评论相关路由
		v1.POST("/comments", commentHandler.CreateComment)
		v1.GET("/posts/:post_id/comments", commentHandler.GetCommentsByPostID)
		v1.PUT("/comments/:id", commentHandler.UpdateComment)
		v1.DELETE("/comments/:id", commentHandler.DeleteComment)
	}
}
