package user

import (
	"getway/consts"
	"getway/service"
	"github.com/gin-gonic/gin"
	"github.com/yanlihongaichila/framework/http"
)

func Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Res(c, 121, nil, err.Error())
		return
	}

	info, err := service.Login(c, loginReq.Username, loginReq.Password)
	if err != nil {
		http.Res(c, 122, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}
