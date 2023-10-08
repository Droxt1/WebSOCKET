package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"server/internal/user"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler) {
	r = gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//swagger:route POST /signup user CreateUserRequest
	// Create a new user
	// responses:
	// 	200: CreateUserResponse
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func Start(addr string) error {
	return r.Run(addr)
}
