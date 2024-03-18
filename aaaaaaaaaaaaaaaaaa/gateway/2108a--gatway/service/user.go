package service

import (
	"context"
	"fmt"
	"gateway/client"
	"gateway/consts"
	"gateway/utils"
	"github.com/JobNing/frameworkJ/redis"
	"github.com/JobNing/message/user"
	"math/rand"
	"time"
)

type LoginRes struct {
	Token    string
	UserInfo *user.UserInfo
}

func Login(ctx context.Context, mobile string, password, code string) (*LoginRes, error) {
	info, err := client.GetUserByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("用户不存在")
	}

	if password != "" {
		if utils.CheckPwd(info.Password, password) {
			return nil, fmt.Errorf("密码错误")
		}
	}

	if code != "" {
		val, err := redis.GetByKey(ctx, consts.ServiceName, consts.SMS_KEY+mobile)
		if err != nil {
			return nil, err
		}
		if val != code {
			return nil, fmt.Errorf("验证码错误")
		}
	}

	token, err := utils.GenToken(info.ID)
	if err != nil {
		return nil, err
	}

	return &LoginRes{
		Token:    token,
		UserInfo: info,
	}, err
}

func RegisterUser(ctx context.Context, mobile string, password, code string) (*LoginRes, error) {
	info, err := client.GetUserByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if info != nil {
		return nil, fmt.Errorf("该手机号已注册")
	}

	val, err := redis.GetByKey(ctx, consts.ServiceName, consts.SMS_KEY+mobile)
	if err != nil {
		return nil, err
	}
	if val != code {
		return nil, fmt.Errorf("验证码错误")
	}

	newPwd, err := utils.GetPwd(password)
	if err != nil {
		return nil, err
	}

	newInfo, err := client.CreateUser(ctx, &user.UserInfo{
		Password: newPwd,
		Mobile:   mobile,
	})
	if err != nil {
		return nil, err

	}

	token, err := utils.GenToken(newInfo.ID)
	if err != nil {
		return nil, err
	}

	return &LoginRes{
		Token:    token,
		UserInfo: newInfo,
	}, err
}

func SendMessage(ctx context.Context, mobile string) error {
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))

	err := redis.SetKey(ctx, consts.ServiceName, consts.SMS_KEY+mobile, code, time.Second*60)
	if err != nil {
		return err
	}
	//发送验证码
	//utils.SendMessage(mobile, fmt.Sprintf("%v", code), "", "")
	return nil
}

func GetUserInfo(ctx context.Context, id int64) (*user.UserInfo, error) {
	info, err := client.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("用户不存在")
	}
	return info, nil
}

func UpdateUserInfo(ctx context.Context, in *user.UserInfo) (*user.UserInfo, error) {
	_, err := GetUserInfo(ctx, in.ID)
	info, err := client.UpdateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return info, nil
}
