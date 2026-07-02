package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	postInfrastructureModel "hobby-blog/internal/post/infrastructure"
	userInfrastructureModel "hobby-blog/internal/user/infrastructure"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err := db.AutoMigrate(&userInfrastructureModel.User{}, &postInfrastructureModel.Post{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db
}
