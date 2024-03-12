package user

import (
	"getway/middleware"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/login", Login)

	u := r.Group("/user")
	u.Use(middleware.AuthMiddleware)
	{
		u.POST("info", GetUserInfo)
	}
}
