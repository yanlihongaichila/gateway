package api

import (
	"getway/api/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	user.Register(r)
}
