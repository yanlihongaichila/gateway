package goods

import (
	"gateway/consts"
	"gateway/service"
	"github.com/JobNing/frameworkJ/http"
	"github.com/gin-gonic/gin"
)

func GetGoodsInfo(c *gin.Context) {
	var req struct {
		GoodID int64 `json:"good_id"`
	}
	if err := c.ShouldBind(&req); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if req.GoodID == 0 {
		http.Res(c, consts.PRM_ERROR, nil, "商品ID不能为空")
		return
	}

	info, err := service.GetGoodInfo(c, req.GoodID)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}

func GetGoodsList(c *gin.Context) {
	var req struct {
		Offset int64 `json:"offset"`
		Limit  int64 `json:"limit"`
		Type   int64 `json:"type"`
	}
	if err := c.ShouldBind(&req); err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}
	if req.Limit == 0 {
		http.Res(c, consts.PRM_ERROR, nil, "Limit不能为0")
		return
	}
	info, err := service.GetGoodList(c, req.Offset, req.Limit, req.Type)
	if err != nil {
		http.Res(c, consts.PRM_ERROR, nil, err.Error())
		return
	}

	http.Res(c, consts.SUCCESS, info, "")
	return
}
