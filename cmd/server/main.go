package main

import (
	"catinder/internal/entity"
	"catinder/internal/handler/v1"
	"catinder/internal/repository"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},                         // 允許的來源
		AllowedMethods:   []string{"GET", "HEAD", "PUT", "PATCH", "POST", "DELETE"}, // 支援的 HTTP 方法
		AllowCredentials: true,                                                      // 允許傳送驗證資訊，例如 Cookie
	})

	// 使用轉換後的 CORS 中間件，將 corsMiddleware.Handler 轉換為 gin.HandlerFunc
	r.Use(func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		c.Next()
	})

	// read .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// get dsn from environment variable
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN environment variable is not set")
	}

	// db connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// migrate user table
	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	// init repository
	repository.InitializeDatabase(db)

	// setup routes
	handler.SetupRoutes(r)

	// run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
