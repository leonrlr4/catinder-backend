package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// 確保你已經導入了 GORM 相關的包
)

func registerUser(c *gin.Context) {
	// 定义接收数据的结构体
	type RegisterInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	db, err := SetupDatabase()
	if err != nil {
		log.Fatal("failed to connect to db error: ", err)
	}
	var regInfo RegisterInfo

	if err := c.ShouldBindJSON(&regInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 对密码进行哈希处理，不要直接保存明文密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(regInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while hashing password"})
		return
	}

	// 创建用户对象
	newUser := User{
		Username: regInfo.Username,
		Password: string(hashedPassword),
	}

	// 将新用户添加到数据库
	result := db.Create(&newUser) // `db` 是 *gorm.DB 类型的数据库连接实例
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func main() {
	r := gin.Default()

	db, err := SetupDatabase()
	if err != nil {
		log.Fatal("failed to connect to db error: ", err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("failed to migrate db error: ", err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello Catinder!"})
	})
	r.POST("/register", registerUser)

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.File("./favicon.ico") // 或者c.File("路徑到你的favicon.ico檔案")
	})

	r.Run() // 在 0.0.0.0:8080 上監聽
}
