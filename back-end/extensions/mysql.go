package extensions

import (
	"MiniShop/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func init() {
	// 加载环境变量
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please create one in the root directory")
	}
}

// GetDB 返回连接成功的数据库
func GetDB() *gorm.DB {
	db, err := gorm.Open(os.Getenv("DB_TYPE"), fmt.Sprintf("%s:%s@(%s)/%s%s", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_OPTIONS")))

	if err != nil {
		panic("连接数据库失败,Error: " + err.Error())
	}

	//自动迁移
	db.AutoMigrate(&models.Good{}, &models.Order{})

	return db
}
