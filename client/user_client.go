package client

import (
	"context"
	"getway/consts"
	"github.com/yanlihongaichila/framework/gprc"
	"github.com/yanlihongaichila/proto/user"
)

type handler func(ctx context.Context, cli user.UserClient) (interface{}, error)

func withClient(ctx context.Context, handler handler) (interface{}, error) {
	conn, err := gprc.Client(consts.SERVICENAME)
	if err != nil {

		return nil, err
	}

	userCli := user.NewUserClient(conn)

	res, err := handler(ctx, userCli)
	if err != nil {
		return nil, err
	}

	conn.Close()
	return res, nil
}

func GetUserByUsername(ctx context.Context, username string) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.GetUserByUsername(ctx, &user.GetUserByUsernameRequest{
			Username: username,
		})
		if err != nil {
			return nil, err
		}
		return res.User, err
	})
	if err != nil {

		return nil, err
	}
	return info.(*user.UserInfo), nil
}

func GetUserById(ctx context.Context, userId int64) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.GetUser(ctx, &user.GetUserRequest{ID: userId})
		if err != nil {
			return nil, err
		}
		return res.User, err
	})
	if err != nil {

		return nil, err
	}
	return info.(*user.UserInfo), nil
}
