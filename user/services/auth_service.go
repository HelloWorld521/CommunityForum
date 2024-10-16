package services

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"gorm.io/gorm"
	"your_project_path/common/models"
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

func RegisterUser(db *gorm.DB, user *models.User) (string, error) {
	// Check if user already exists
	var existingUser models.User
	if db.Where("email = ?", user.Email).First(&existingUser).RecordNotFound() {
		// Hash the password
		hashedPassword, err := HashPassword(user.PasswordHash)
		if err != nil {
			return "", err
		}
		user.PasswordHash = hashedPassword

		// Save the user to the database
		if err := db.Create(user).Error; err != nil {
			return "", err
		}

		return "User registered successfully", nil
	} else {
		return "", fmt.Errorf("user with this email already exists")
	}
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
