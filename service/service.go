package service

import (
	"context"
	"fmt"
	"getway/client"
	"github.com/yanlihongaichila/proto/user"
)

type LoginRes struct {
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

	return &LoginRes{
		UserInfo: info,
	}, err
}
