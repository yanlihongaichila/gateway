package client

import (
	"context"
	"github.com/JobNing/frameworkJ/grpc"
	"github.com/JobNing/message/goods"
	//"goods/consts"
)

type goodsHandler func(ctx context.Context, cli goods.GoodClient) (interface{}, error)

func withGoodClient(ctx context.Context, handler goodsHandler) (interface{}, error) {
	conn, err := grpc.Client(ctx, "DEFAULT_GROUP", "goods")
	if err != nil {
		return nil, err
	}

	goodsCli := goods.NewGoodClient(conn)

	res, err := handler(ctx, goodsCli)
	if err != nil {
		return nil, err
	}

	conn.Close()
	return res, nil
}

func GetGood(ctx context.Context, id int64) (*goods.GoodInfo, error) {
	info, err := withGoodClient(ctx, func(ctx context.Context, cli goods.GoodClient) (interface{}, error) {
		res, err := cli.GetGood(ctx, &goods.GetGoodRequest{ID: id})
		if err != nil {
			return nil, err
		}
		return res.Info, err
	})
	if err != nil {
		return nil, err
	}
	return info.(*goods.GoodInfo), nil
}

func GetGoodList(ctx context.Context, offset, limit, goodType int64) ([]*goods.GoodInfo, error) {
	infos, err := withGoodClient(ctx, func(ctx context.Context, cli goods.GoodClient) (interface{}, error) {
		res, err := cli.GetGoods(ctx, &goods.GetGoodsRequest{
			Offset: offset,
			Limit:  limit,
			Type:   goodType,
		})
		if err != nil {
			return nil, err
		}
		return res.Infos, err
	})
	if err != nil {
		return nil, err
	}
	return infos.([]*goods.GoodInfo), nil
}
