package main

import (
	"github.com/gin-gonic/gin"

	"hobby-blog/internal/db"
	"hobby-blog/internal/router"
	"hobby-blog/internal/handler"
)

func main()  {
	dbConn := db.ConnectDB()

	handlers := handler.NewHandlers(dbConn)

	r := router.SetUpRouter(handlers)

	r.Run(":8080")
}
