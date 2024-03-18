package api

import (
	"gateway/api/goods"
	"gateway/api/user"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	user.Register(r)
	goods.Register(r)
}
