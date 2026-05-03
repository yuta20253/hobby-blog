package main

import (
	"hobby-blog/internal/db"
	"hobby-blog/internal/router"
	"hobby-blog/internal/container"
)

func main()  {
	dbConn := db.ConnectDB()

	c := container.NewContainer(dbConn)

	r := router.SetUpRouter(c.AuthHandler)

	r.Run(":8080")
}
