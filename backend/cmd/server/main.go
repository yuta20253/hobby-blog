package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"

	"hobby-blog/internal/db"
	"hobby-blog/internal/models"
)

func main()  {
	r := gin.Default()
	r.Use(cors.Default())

	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err := dbConn.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Migration failed:", err)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
