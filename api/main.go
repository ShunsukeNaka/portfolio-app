package main

import (
	"api/handlers"
	"api/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()

	err := models.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if err != nil {
		log.Fatalf("Failed to enable UUID extension: %v", err)
	}

	err = models.DB.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Tag{},
		&models.Comment{},
		&models.Follow{},
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

	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.GET("/", handlers.GetUsers)
	}

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
