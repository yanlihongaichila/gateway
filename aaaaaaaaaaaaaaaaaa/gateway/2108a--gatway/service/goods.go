package service

import (
	"context"
	"fmt"
	"gateway/client"
	"github.com/JobNing/message/goods"
)

func GetGoodInfo(ctx context.Context, id int64) (*goods.GoodInfo, error) {
	info, err := client.GetGood(ctx, id)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("用户不存在")
	}
	return info, nil
}

func GetGoodList(ctx context.Context, offset, limit, goodType int64) ([]*goods.GoodInfo, error) {
	info, err := client.GetGoodList(ctx, offset, limit, goodType)
	if err != nil {
		return nil, err
	}
	return info, nil
}
