package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/model/trans"
	"github.com/ziyoubiancheng/goshop/router/mdw"
	"go.uber.org/zap"
)

type plugin struct{}

func InitPlugin() *plugin {
	return &plugin{}
}

// Create 新增一条记录
func (*plugin) Create(c *gin.Context, db *gorm.DB, data *mysql.Plugin) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create plugin create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*plugin) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("plugin").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("plugin update error", zap.String("err", err.Error()))
		return
	}
	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*plugin) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("plugin").Where(sql, binds...).Delete(&mysql.Plugin{}).Error; err != nil {
		model.Logger.Error("plugin delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*plugin) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.Plugin, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("plugin").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("plugin info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*plugin) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.Plugin, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("plugin").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("plugin info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListPage 根据分页条件查询list
func (*plugin) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.Plugin) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("plugin").Where(sql, binds...)
	respList = make([]mysql.Plugin, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
