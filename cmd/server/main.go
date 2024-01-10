package main

import (
	"catinder/internal/handler"
	"catinder/internal/model"
	"catinder/internal/repository"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	// db connect
	dsn := "root:password@tcp(127.0.0.1:3306)/catinder?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// migrate user table
	if err := db.AutoMigrate(&model.User{}); err != nil {
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
