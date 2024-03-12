package service

import (
	"context"
	"fmt"
	"getway/client"
	"getway/utils"
	"github.com/yanlihongaichila/proto/user"
)

type LoginRes struct {
	Token    string         `json:"token"`
	UserInfo *user.UserInfo `json:"user_info"`
}

func Login(ctx context.Context, username string, password string) (*LoginRes, error) {
	info, err := client.GetUserByUsername(ctx, username)
	if err != nil {

		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("用户不存在")
	}

	if info.Password != password {
		return nil, fmt.Errorf("密码错误")
	}

	//生成token
	token, err := utils.GenToken(info.ID)
	if err != nil {
		return nil, err
	}
	return &LoginRes{
		Token:    token,
		UserInfo: info,
	}, err
}

func GetUser(ctx context.Context, userId int64) (*user.UserInfo, error) {
	info, err := client.GetUserById(ctx, userId)
	if err != nil {
		return nil, nil
	}

	if info == nil {
		return nil, fmt.Errorf("用户不存在")
	}
	return info, nil
}
