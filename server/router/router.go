package router

import (
	"github.com/gin-gonic/gin"
	"server/internal/user"
	"server/internal/ws"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()
	v1Routes := r.Group("/api/v1")
	{
		userRoutes := v1Routes.Group("/user")
		{
			userRoutes.POST("/signup", userHandler.CreateUser)
			userRoutes.POST("/login", userHandler.Login)
			userRoutes.GET("/logout", userHandler.Logout)

		}

		wsRoutes := v1Routes.Group("/ws")
		{
			wsRoutes.POST("/createRoom", wsHandler.CreateRoom)
			wsRoutes.GET("/joinRoom/:roomId", wsHandler.JoinRoom)
			wsRoutes.GET("/getRooms", wsHandler.GetRooms)
			wsRoutes.GET("/get/:roomId", wsHandler.GetClients)

		}
	}
}

func Start(addr string) error {
	return r.Run(addr)
}
