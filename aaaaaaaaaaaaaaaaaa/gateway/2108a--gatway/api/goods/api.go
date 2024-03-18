package goods

import (
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	u := r.Group("/goods")
	{
		u.POST("info", GetGoodsInfo)
		u.POST("list", GetGoodsList)
	}
}
