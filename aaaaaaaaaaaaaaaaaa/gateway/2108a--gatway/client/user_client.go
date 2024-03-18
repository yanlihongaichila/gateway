package client

import (
	"context"
	"github.com/JobNing/frameworkJ/grpc"
	"github.com/JobNing/message/user"
	//"user/consts"
)

type handler func(ctx context.Context, cli user.UserClient) (interface{}, error)

func withClient(ctx context.Context, handler handler) (interface{}, error) {
	conn, err := grpc.Client(ctx, "DEFAULT_GROUP", "user")
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

func GetUser(ctx context.Context, id int64) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.GetUser(ctx, &user.GetUserRequest{ID: id})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}

func GetUserByUsername(ctx context.Context, username string) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.GetUserByUsername(ctx, &user.GetUserByUsernameRequest{
			Username: username,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}

func GetUserByMobile(ctx context.Context, mobile string) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.GetUserByMobile(ctx, &user.GetUserByMobileRequest{
			Mobile: mobile,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}

func CreateUser(ctx context.Context, in *user.UserInfo) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.CreateUser(ctx, &user.CreateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}

func UpdateUser(ctx context.Context, in *user.UserInfo) (*user.UserInfo, error) {
	info, err := withClient(ctx, func(ctx context.Context, cli user.UserClient) (interface{}, error) {
		res, err := cli.UpdateUser(ctx, &user.UpdateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*user.UserInfo), nil
}
