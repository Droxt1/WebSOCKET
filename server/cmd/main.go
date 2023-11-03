package main

import (
	_ "github.com/lib/pq"
	"log"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
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

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)

	err = router.Start(":3000")
	if err != nil {
		return
	}

}
