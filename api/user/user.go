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

func GetUserInfo(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(int64)

	info, err := service.GetUser(c, userID)
	if err != nil {
		http.Res(c, consts.PPM_ERROR, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}
