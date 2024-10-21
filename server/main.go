package main

import (
	"community/user/handlers"
	"community/user/models"
	"community/user/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	dsn := "root:123456@tcp(192.168.3.3:13306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// 自动迁移数据库模式
	db.AutoMigrate(&models.User{})

	// 初始化 AuthService
	authService := &services.AuthService{DB: db}

	// 创建 Gin 路由
	r := gin.Default()
	r.Use(corsMiddleware())

	// 初始化 AuthHandler 并设置路由
	authHandler := &handlers.AuthHandler{AuthService: authService}
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authHandler.Register)
		v1.POST("/login", authHandler.Login)
	}

	// 启动服务器
	if err := r.Run(":8081"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// corsMiddleware 设置CORS中间件
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
