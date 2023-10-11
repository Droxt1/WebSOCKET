package main

import (
	_ "github.com/lib/pq"
	"log"
	_ "server/cmd/docs"
	"server/db"
	"server/internal/user"
	"server/internal/ws"
	"server/router"
)

//	@title			Websocket API
//	@description	This is a websocket server for chat application
//	@version		1
//	@host			localhost:8080
//	@BasePath		/api/v1
//swagger:route POST /signup user CreateUserRequest
// Create a new user
// responses:git push -u origin main
// 	200: CreateUserResponse

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

	router.Start("0.0.0.0:8080")

}
