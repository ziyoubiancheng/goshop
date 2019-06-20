package wechat

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/service"
)

var DefaultWechatUid = "github.com/ziyoubiancheng/goshop/app/middleware/wechatUser"
var DefaultWechatUsername = "github.com/ziyoubiancheng/goshop/app/middleware/wechatUsername"

func WechatAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := c.GetHeader("Access-Token")
		if !service.AccessToken.CheckAccessToken(accessToken) {
			c.JSON(http.StatusMethodNotAllowed, nil)
			c.Abort()
		}

		userInfo, err := service.AccessToken.DecodeAccessToken(accessToken)
		if err != nil {
			c.JSON(http.StatusMethodNotAllowed, nil)
			c.Abort()
		}
		uid, flag := userInfo["sub"].(float64)
		if !flag {
			c.JSON(http.StatusMethodNotAllowed, nil)
			c.Abort()
		}
		c.Set(DefaultWechatUid, int(uid))
		c.Next()
	}
}
func WechatUid(c *gin.Context) int {
	return c.MustGet(DefaultWechatUid).(int)
}

func WechatUserName(c *gin.Context) (username string) {
	value, flag := c.Get(DefaultWechatUsername)
	if flag {
		username = value.(string)
		return
	}

	uid := c.MustGet(DefaultWechatUid).(int)
	user := mysql.UserOpen{}
	err := model.Db.Where("uid = ?", uid).Find(&user)
	if err != nil {
		// todo log
		return
	}
	c.Set(DefaultWechatUsername, user.Nickname)
	username = user.Nickname
	return
}
