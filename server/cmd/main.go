package main

import (
	_ "github.com/lib/pq"
	"log"
	_ "server/cmd/docs"
	"server/db"
	"server/internal/user"
	"server/router"
)

//	@title			Websocket API
//	@description	This is a websocket server for chat application
//	@version		1
//	@host			localhost:8080
//	@BasePath		/api/v1
//swagger:route POST /signup user CreateUserRequest
// Create a new user
// responses:
// 	200: CreateUserResponse

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
