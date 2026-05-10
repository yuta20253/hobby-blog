package main

import (
	"hobby-blog/internal/container"
	"hobby-blog/internal/db"
	"hobby-blog/internal/router"
	"log"
)

func main() {
	dbConn := db.ConnectDB()

	c := container.NewContainer(dbConn)

	r := router.SetUpRouter(
		c.AuthHandler,
		c.PostHandler,
		c.MypageHandler,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
