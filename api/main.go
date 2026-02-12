package main

import (
	"fmt"
	"log"

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

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
