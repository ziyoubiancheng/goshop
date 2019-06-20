package shop

import (
	"github.com/ziyoubiancheng/goshop/model/resp"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/router/base"
	"github.com/ziyoubiancheng/goshop/router/mdw"
	"github.com/ziyoubiancheng/goshop/service"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Info(c *gin.Context) {
	info, _ := service.Shop.InfoX(c, mysql.Conds{
		"open_id": mdw.OpenId(c),
	})

	base.JSON(c, base.MsgOk, resp.ShopInfo{
		Name:                         info.Name,
		Logo:                         info.Logo,
		ContactNumber:                info.ContactNumber,
		Description:                  info.Description,
		ColorScheme:                  info.ColorScheme,
		PortalTemplateId:             info.PortalTemplateId,
		WechatPlatformQr:             info.WechatPlatformQr,
		GoodsCategoryStyle:           info.GoodsCategoryStyle,
		Host:                         info.Host,
		OrderAutoCloseExpires:        info.OrderAutoCloseExpires,
		OrderAutoConfirmExpires:      info.OrderAutoConfirmExpires,
		OrderAutoCloseRefoundExpires: info.OrderAutoCloseRefoundExpires,
	})
}
