package goods

import (
	"github.com/gin-gonic/gin"

	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/model/trans"
	"github.com/ziyoubiancheng/goshop/pkg/tree"
	"github.com/ziyoubiancheng/goshop/router/base"
	"github.com/ziyoubiancheng/goshop/router/mdw"
	"github.com/ziyoubiancheng/goshop/service"
)

func CategoryList(c *gin.Context) {
	list, _ := service.GoodsCategory.List(c, mysql.Conds{
		"open_id": mdw.OpenId(c),
	}, "update_time desc")

	respTree := tree.NewTreeStore(list).ListToTree(0)

	base.JSONList(c, respTree, 0)

}

func CategoryInfo(c *gin.Context) {
	req := trans.ReqGoodscategoryInfo{}
	if err := c.Bind(&req); err != nil {
		base.JSON(c, base.MsgErr)
		return
	}

	info, err := service.GoodsCategory.InfoX(c, mysql.Conds{
		"id":      req.Id,
		"open_id": mdw.OpenId(c),
	})
	if err != nil {
		base.JSON(c, base.MsgErr)
		return
	}
	base.JSON(c, base.MsgOk, map[string]interface{}{
		"info": info,
	})

}
