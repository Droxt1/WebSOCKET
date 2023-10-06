package main

import (
	_ "github.com/lib/pq"
	"log"
	"server/db"
	"server/internal/user"
	"server/router"
)

func main() {
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}
	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	router.InitRouter(userHandler)
	if err := router.Start(":8080"); err != nil {
		log.Fatalf("could not start server: %v", err)
	}

}
