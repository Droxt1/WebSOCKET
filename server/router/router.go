package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/user"
)

// Engine is the framework's instance.
// it contains the muxer, middleware and configuration settings.
// Create an instance of Engine by using New() or Default().
// The Engine is the primary interface for the users to interact with the framework.
var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()
	r.POST("/signup", userHandler.CreateUser)

}

func Start(addr string) error {
	return r.Run(addr)

}
