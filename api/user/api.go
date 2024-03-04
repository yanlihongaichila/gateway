package user

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.POST("/login", Login)

}
