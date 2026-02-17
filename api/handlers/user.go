package handlers

import (
	"api/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest

	// リクエストボディのJSON -> req
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なリクエストです"})
		return
	}

	// DBからユーザー検索
	var user models.User
	if err := models.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが正しくありません。"})
		return
	}

	// パスワードの照合
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスまたはパスワードが正しくありません。"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv(("JWT_SECRET"))))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの作成に失敗しました。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user":  user,
	})
}

// ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓ トークンチェック有 ↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓↓
func GetMyProfile(c *gin.Context) {
	id, exists := c.Get("userID")

	if !exitsts {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "認証情報を確認できません"})
		return
	}

	userID := id.(uuid.UUID)

	// Articlesも同時に取得
	var user models.User
	err := models.DB.Preload("Articles", func(db *gorm.DB) *gorm.DB{
		return db.Select("title", "user_id")
	}).First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}
	
	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	var users []models.User
	// 関連する記事(Articles)も一緒に取得
	models.DB.Preload("Articles").Find(&users)
	c.JSON(http.StatusOK, users)
}
