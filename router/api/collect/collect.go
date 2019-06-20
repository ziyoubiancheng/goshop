package collect

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/model/trans"
	"github.com/ziyoubiancheng/goshop/router/base"
	"github.com/ziyoubiancheng/goshop/router/mdw"
	"github.com/ziyoubiancheng/goshop/router/mdw/wechat"
	"github.com/ziyoubiancheng/goshop/service"
)

func Info(c *gin.Context) {
	uid := wechat.WechatUid(c)
	_, respArr := service.GoodsCollect.ListPage(c, mysql.Conds{
		"uid": uid,
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
		Sort:     "id asc",
	})

	goodsIs := make([]int, 0)
	for _, value := range respArr {
		goodsIs = append(goodsIs, value.GoodsId)
	}

	cnt, output := service.Goods.ListPage(c, mysql.Conds{
		"id": mysql.Cond{"in", goodsIs},
	}, trans.ReqPage{
		Current:  0,
		PageSize: 1000,
		Sort:     "id asc",
	})
	base.JSONList(c, output, cnt)
}

func Create(c *gin.Context) {
	req := trans.ReqCollectCreate{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	openId := mdw.OpenId(c)
	var err error
	tx := model.Db.Begin()
	err = service.GoodsCollect.DeleteX(c, tx, mysql.Conds{
		"uid":      uid,
		"goods_id": req.GoodsId,
	})

	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	_, err = service.Goods.InfoX(c, mysql.Conds{
		"id": req.GoodsId,
	})
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	createInfo := &mysql.GoodsCollect{
		Uid:        uid,
		GoodsId:    req.GoodsId,
		CreateTime: time.Now().Unix(),
		OpenId:     openId,
	}
	err = service.GoodsCollect.Create(c, tx, createInfo)
	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}

	base.JSON(c, base.MsgOk)
}

func Del(c *gin.Context) {
	req := trans.ReqCollectDel{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	uid := wechat.WechatUid(c)
	var err error
	tx := model.Db.Begin()
	err = service.GoodsCollect.DeleteX(c, tx, mysql.Conds{
		"uid":      uid,
		"goods_id": req.GoodsId,
	})

	if err != nil {
		tx.Rollback()
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk)
}
