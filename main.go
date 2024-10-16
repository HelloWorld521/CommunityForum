package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"user/handlers"
	"user/models"
)

func main() {
	dsn := "community_user:your_password@tcp(db:3306)/community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	authService := &handlers.AuthService{DB: db}

	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", authService.Register)
		v1.POST("/login", authService.Login)
	}

	r.Run(":8081")
}
