package user

import (
	"gateway/consts"
	"gateway/service"
	"github.com/JobNing/frameworkJ/http"
	"github.com/JobNing/message/user"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(int64)

	info, err := service.GetUserInfo(c, userID)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}

func UpdateUserInfo(c *gin.Context) {
	uid, _ := c.Get("user_id")
	userID := uid.(int64)

	var req = new(user.UserInfo)
	if err := c.ShouldBind(&req); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	req.ID = userID

	info, err := service.UpdateUserInfo(c, req)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}
