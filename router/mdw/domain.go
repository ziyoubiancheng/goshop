package mdw

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/model/mysql"
)

var DefaultDomainId = "github.com/ziyoubiancheng/goshop/app/middleware/domainId"
var DefaultDomainName = "github.com/ziyoubiancheng/goshop/app/middleware/domainName"

func Domain() gin.HandlerFunc {
	return func(c *gin.Context) {
		domain := c.Param("domain")
		var openUser mysql.Biz
		model.Db.Select("open_id,domain").Where("domain=?", domain).Find(&openUser)
		fmt.Println(openUser)
		if openUser.OpenId == 0 {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
			return
		}

		c.Set(DefaultDomainId, openUser.OpenId)
		c.Set(DefaultDomainName, domain)
		c.Next()
	}
}

func OpenId(c *gin.Context) int {
	return c.MustGet(DefaultDomainId).(int)
}

func DomainName(c *gin.Context) string {
	return c.MustGet(DefaultDomainName).(string)
}
