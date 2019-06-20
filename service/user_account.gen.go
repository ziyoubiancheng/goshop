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

type userAccount struct{}

func InitUserAccount() *userAccount {
	return &userAccount{}
}

// Create 新增一条记录
func (*userAccount) Create(c *gin.Context, db *gorm.DB, data *mysql.UserAccount) (err error) {
	data.OpenId = mdw.OpenId(c)

	if err = db.Create(data).Error; err != nil {
		model.Logger.Error("create userAccount create error", zap.String("err", err.Error()))
		return
	}
	return nil
}

// Update 根据主键更新一条记录
func (*userAccount) Update(c *gin.Context, db *gorm.DB, paramUid int, ups mysql.Ups) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramUid}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_account").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_account update error", zap.String("err", err.Error()))
		return
	}
	return
}

// UpdateX Update的扩展方法，根据Cond更新一条或多条记录
func (*userAccount) UpdateX(c *gin.Context, db *gorm.DB, conds mysql.Conds, ups mysql.Ups) (err error) {

	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_account").Where(sql, binds...).Updates(ups).Error; err != nil {
		model.Logger.Error("user_account update error", zap.String("err", err.Error()))
		return
	}
	return
}

// Delete 根据主键删除一条记录。如果有delete_time则软删除，否则硬删除。
func (*userAccount) Delete(c *gin.Context, db *gorm.DB, paramUid int) (err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramUid}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_account").Where(sql, binds...).Delete(&mysql.UserAccount{}).Error; err != nil {
		model.Logger.Error("user_account delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// DeleteX Delete的扩展方法，根据Cond删除一条或多条记录。如果有delete_time则软删除，否则硬删除。
func (*userAccount) DeleteX(c *gin.Context, db *gorm.DB, conds mysql.Conds) (err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = db.Table("user_account").Where(sql, binds...).Delete(&mysql.UserAccount{}).Error; err != nil {
		model.Logger.Error("user_account delete error", zap.String("err", err.Error()))
		return
	}

	return
}

// Info 根据PRI查询单条记录
func (*userAccount) Info(c *gin.Context, paramUid int) (resp mysql.UserAccount, err error) {
	var sql = "`uid`=?"
	var binds = []interface{}{paramUid}
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("user_account").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_account info error", zap.String("err", err.Error()))
		return
	}
	return
}

// InfoX Info的扩展方法，根据Cond查询单条记录
func (*userAccount) InfoX(c *gin.Context, conds mysql.Conds) (resp mysql.UserAccount, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	if err = model.Db.Table("user_account").Where(sql, binds...).First(&resp).Error; err != nil {
		model.Logger.Error("user_account info error", zap.String("err", err.Error()))
		return
	}
	return
}

// List 查询list，extra[0]为sorts
func (*userAccount) List(c *gin.Context, conds mysql.Conds, extra ...string) (resp []mysql.UserAccount, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	sorts := ""
	if len(extra) >= 1 {
		sorts = extra[0]
	}
	if err = model.Db.Table("user_account").Where(sql, binds...).Order(sorts).Find(&resp).Error; err != nil {
		model.Logger.Error("user_account info error", zap.String("err", err.Error()))
		return
	}
	return
}

// ListMap 查询map，map遍历的时候是无序的，所以指定sorts参数没有意义
func (*userAccount) ListMap(c *gin.Context, conds mysql.Conds) (resp map[int]mysql.UserAccount, err error) {
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	mysqlSlice := make([]mysql.UserAccount, 0)
	resp = make(map[int]mysql.UserAccount, 0)
	if err = model.Db.Table("user_account").Where(sql, binds...).Find(&mysqlSlice).Error; err != nil {
		model.Logger.Error("user_account info error", zap.String("err", err.Error()))
		return
	}
	for _, value := range mysqlSlice {
		resp[value.Uid] = value
	}
	return
}

// ListPage 根据分页条件查询list
func (*userAccount) ListPage(c *gin.Context, conds mysql.Conds, reqList trans.ReqPage) (total int, respList []mysql.UserAccount) {
	if reqList.PageSize == 0 {
		reqList.PageSize = 10
	}
	if reqList.Current == 0 {
		reqList.Current = 1
	}
	sql, binds := mysql.BuildQuery(conds)
	sql += " and open_id=?"
	binds = append(binds, mdw.OpenId(c))

	db := model.Db.Table("user_account").Where(sql, binds...)
	respList = make([]mysql.UserAccount, 0)
	db.Count(&total)
	db.Order(reqList.Sort).Offset((reqList.Current - 1) * reqList.PageSize).Limit(reqList.PageSize).Find(&respList)
	return
}
