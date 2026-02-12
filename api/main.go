package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDB()

	err := DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Fatalf("Failed to enable UUID extension: %v", err)
	}

	err = DB.AutoMigrate(
		&User{},
		&Article{},
		&Tag{},
		&Comment{},
		&Follow{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	fmt.Println("Migration Copleted")

	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートURL ("/") に対するGETリクエストをハンドル
	r.GET("/", func(c *gin.Context) {
		// JSONレスポンスを返す
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.POST("/api/users", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := DB.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, user)
	})

	r.GET("/api/users", func(c *gin.Context) {
		var users []User
		// 関連する記事(Articles)も一緒に取得
		DB.Preload("Articles").Find(&users)
		c.JSON(http.StatusOK, users)
	})

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
