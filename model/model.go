package model

import (
	"github.com/jinzhu/gorm"
	"github.com/ziyoubiancheng/drivers/pkg/cache/redis"
	"github.com/ziyoubiancheng/drivers/pkg/database/mysql"
	"github.com/ziyoubiancheng/drivers/pkg/logger"
)

var (
	Logger *logger.Client
	Db     *gorm.DB
	Redigo *redis.Client
)

func Init() {
	Db = mysql.Caller("mall")
	Logger = logger.Caller("system")
	Redigo = redis.Caller("auth")
}
