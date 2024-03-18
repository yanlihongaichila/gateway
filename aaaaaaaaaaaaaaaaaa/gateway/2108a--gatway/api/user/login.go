package user

import (
	"gateway/consts"
	"gateway/service"
	"github.com/JobNing/frameworkJ/http"
	"github.com/gin-gonic/gin"
	"regexp"
)

func Login(c *gin.Context) {
	var loginReq struct {
		Mobile   string `json:"mobile"`
		Password string `json:"password"`
		Code     string `json:"code"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if loginReq.Mobile == "" {
		http.Res(c, consts.PRM_ERROR, nil, "手机号不能为空")
		return
	}
	if !checkPhone(loginReq.Mobile) {
		http.Res(c, consts.PRM_ERROR, nil, "手机号格式不正确")
		return
	}
	if loginReq.Password == "" && loginReq.Code == "" {
		http.Res(c, consts.PRM_ERROR, nil, "密码或验证码不能为空")
		return
	}

	info, err := service.Login(c, loginReq.Mobile, loginReq.Password, loginReq.Code)
	if err != nil {
		http.Res(c, consts.SYSTEM_ERROR, nil, err.Error())
		return
	}
	if info == nil {
		http.Res(c, consts.NOTFOUND, "", "账号不存在")
	}
	http.Res(c, consts.SUCCESS, info, "")
	return
}

func RegisterUser(c *gin.Context) {
	var loginReq struct {
		Mobile   string `json:"mobile"`
		Code     string `json:"code"`
		Password string `json:"password"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if !checkPhone(loginReq.Mobile) {
		http.Res(c, consts.PRM_ERROR, nil, "手机号格式不正确")
		return
	}
	if loginReq.Code == "" || loginReq.Mobile == "" || loginReq.Password == "" {
		http.Res(c, consts.PRM_ERROR, nil, "请求参数错误")
		return
	}

	info, err := service.RegisterUser(c, loginReq.Mobile, loginReq.Password, loginReq.Code)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if info == nil {
		http.Res(c, consts.NOTFOUND, "", "账号不存在")
	}
	http.Res(c, consts.SUCCESS, info, "")
	return
}

func checkPhone(phone string) bool {
	regRuler := "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regRuler)
	return reg.MatchString(phone)
}

func SendMessage(c *gin.Context) {
	var loginReq struct {
		Mobile string `json:"mobile"`
	}
	if err := c.ShouldBind(&loginReq); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if !checkPhone(loginReq.Mobile) {
		http.Res(c, consts.PRM_ERROR, nil, "手机号格式不正确")
		return
	}
	err := service.SendMessage(c, loginReq.Mobile)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	http.Res(c, consts.SUCCESS, nil, "")
	return
}
