package main

import (
	"hobby-blog/internal/config"
	"hobby-blog/internal/container"
	"hobby-blog/internal/db"
	"hobby-blog/internal/router"
	"log"
)

func main() {
	dbConn := db.ConnectDB()

	c := container.NewContainer(dbConn)

	cfg := config.Load()

	r := router.SetUpRouter(
		cfg,
		c.AuthHandler,
		c.PostHandler,
		c.MypageHandler,
		c.MediaHandler,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
