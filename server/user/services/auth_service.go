package services

import (
	"community/user/models"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var jwtKey = []byte("your_secret_key")

func HashPassword(password string) (string, error) {
	hash := sha256.New()
	_, err := hash.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(hash.Sum(nil)), nil
}

func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userID,
	})
	return token.SignedString(jwtKey)
}

// RegisterUser 注册新用户
func RegisterUser(db *gorm.DB, user *models.User) (string, error) {
	// 检查是否已经存在具有相同电子邮件的用户
	var existingUser models.User
	result := db.Where("email = ?", user.Email).First(&existingUser)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 对密码进行哈希处理
		hashedPassword, err := HashPassword(user.PasswordHash)
		if err != nil {
			return "", fmt.Errorf("failed to hash password: %w", err)
		}
		user.PasswordHash = hashedPassword

		// 将新用户保存到数据库
		if err := db.Create(user).Error; err != nil {
			return "", fmt.Errorf("failed to create user: %w", err)
		}

		return "User registered successfully", nil
	} else if result.Error != nil {
		// 如果有其他类型的错误
		return "", fmt.Errorf("database query failed: %w", result.Error)
	}

	// 如果找到了具有相同电子邮件的用户
	return "", fmt.Errorf("user with email %s already exists", user.Email)
}

func LoginUser(db *gorm.DB, email, password string) (string, error) {
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Verify the password
	hashedPassword, _ := HashPassword(password)
	if user.PasswordHash != hashedPassword {
		return "", fmt.Errorf("invalid email or password")
	}

	// Generate and return a JWT token
	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}
