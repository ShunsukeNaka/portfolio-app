package main

import (
	"api/handlers"
	"api/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
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

	// トークンをチェックしない
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.GET("/", handlers.GetUsers)
	}

	//　以下はトークンをチェックする
	protectedUserRoutes := r.Group("/users")
	protectedUserRoutes.Use(AuthMiddleware())
	{
		protectedUserRoutes.GET("/getme", handlers.GetMyProfile)
	}

	// articleRoutes := r.Group("/article")
	// {
	// }
	// likeRoutes := r.Group("/likes")
	// {
	// }
	// commentRoutes := r.Group("/comments")
	// {
	// }
	// shareRoutes := r.Group("/shares")
	// {
	// }

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}

func ParseToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 署名方法の確認
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return uuid.Nil, err
	}

	// 取り出す情報を指定
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDStr := claims["user_id"].(string)
		return uuid.Parse(userIDStr)
	}

	return uuid.Nil, fmt.Errorf("invalid token claims")
}

// =================ミドルウェアの定義==========================
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "認証トークンが必要です"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "無効なトークンです"})
			c.Abort()
			return
		}

		// 3. コンテキスト(c)にユーザーIDをセットしておく
		// これで、後のハンドラ関数で userID を取り出せるようになる
		c.Set("userID", userID)
		c.Next()
	}
}

//=================ミドルウェアの定義=↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑↑=
