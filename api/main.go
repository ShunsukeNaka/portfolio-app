package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
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
