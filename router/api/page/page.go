package page

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/model/trans"
	"github.com/ziyoubiancheng/goshop/router/base"
	"github.com/ziyoubiancheng/goshop/service"
)

func Portal(c *gin.Context) {
	value, err := service.Page.InfoX(c, mysql.Conds{
		"is_portal": 1,
	})

	if err != nil {
		base.JSON(c, base.MsgErr, "portal is error")
		return
	}
	resp := trans.Page{
		Id:              value.Id,
		Name:            value.Name,
		Description:     value.Description,
		Body:            json.RawMessage(value.Body),
		IsPortal:        value.IsPortal,
		IsSystem:        value.IsSystem,
		BackgroundColor: value.BackgroundColor,
		Type:            value.Type,
		CreateTime:      value.CreateTime,
		UpdateTime:      value.UpdateTime,
		Module:          value.Module,
		CloneFromId:     value.CloneFromId,
	}
	base.JSON(c, base.MsgOk, resp)
}
