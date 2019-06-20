package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ziyoubiancheng/goshop/model"
	"github.com/ziyoubiancheng/goshop/model/mysql"
	"github.com/ziyoubiancheng/goshop/model/trans"
	"github.com/ziyoubiancheng/goshop/router/mdw"
	"go.uber.org/zap"
)

type coupon struct{}

func InitCoupon() *coupon {
	return &coupon{}
}

// Create 新增一条记录
func (*coupon) Create(c *gin.Context, db *gorm.DB, data *mysql.Coupon) (err error) {
	data.OpenId = mdw.OpenId(c)
	data.CreateTime = time.Now().Unix()
	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create coupon create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*coupon) Update(c *gin.Context, db *gorm.DB, paramId int, ups mysql.Ups) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("coupon").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("coupon update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*coupon) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("coupon").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("coupon update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*coupon) Delete(c *gin.Context, db *gorm.DB, paramId int) (err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("coupon").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("coupon delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*coupon) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("coupon").Where(sql, binds...).Update(mysql.Ups{"delete_time": time.Now().Unix()}).Error; err != nil {
		model.Logger.Error("coupon delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*coupon) Info(c *gin.Context, paramId int) (resp mysql.Coupon, err error) {
	var sql = "`id`=?"
	var binds = []interface{}{paramId}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	if err = model.Db.Table("coupon").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("coupon info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*coupon) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.Coupon, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	if err = model.Db.Table("coupon").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("coupon info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*coupon) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.Coupon, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("coupon").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("coupon info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*coupon) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.Coupon, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	mysqlSlice := make([]mysql.Coupon, 0)
	resp = make(map[int]mysql.Coupon, 0)
	if err = model.Db.Table("coupon").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("coupon info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Id] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*coupon) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.Coupon) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sql += " AND delete_time = 0"
	db := model.Db.Table("coupon").Where(sql, binds...)
	respList = make([]mysql.Coupon, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
