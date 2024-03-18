package user

import (
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/login", Login)
	r.POST("/register", RegisterUser)
	r.POST("/send/message", SendMessage)

	u := r.Group("/user")
	u.Use(middleware.AuthMiddleware)
	{
		u.POST("info", GetUserInfo)
		u.POST("update", UpdateUserInfo)
	}
}
