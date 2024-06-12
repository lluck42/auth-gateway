package admin

import (
	"encoding/json"
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model/admin"

	"github.com/gin-gonic/gin"
)

func RoleAllList(c *gin.Context) {

	var list []struct {
		ID   uint
		Name string
	}

	var tx = bean.MysqlDb.Model(&admin.AdminRole{})

	// 查询列表
	if err := tx.Model(&admin.AdminRole{}).Order("`id` ASC").Find(&list).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, gin.H{
		"list": list,
	})
}

func RoleList(c *gin.Context) {

	var req struct {
		Page  int    `binding:"required,gte=1"`
		Limit int    `binding:"required,gte=1"`
		Name  string `binding:"max=255"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var offset = (req.Page - 1) * req.Limit

	// 查询
	var count int64
	var list []admin.AdminRole

	var tx = bean.MysqlDb.Model(&admin.AdminRole{})

	if req.Name != "" {
		tx = tx.Where("name like ?", "%"+req.Name+"%")
	}

	// 查询总数
	if err := tx.Count(&count).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	// 查询列表
	if err := tx.Limit(req.Limit).Offset(offset).Order("id DESC").Find(&list).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, gin.H{
		"count": count,
		"list":  list,
	})
}

func RoleDetail(c *gin.Context) {
	// 查询列表
	var req struct {
		ID int `binding:"required,gte=1"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var one admin.AdminRole

	if err := bean.MysqlDb.Where("id=?", req.ID).First(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, one)
}

func RoleEdit(c *gin.Context) {
	// 查询列表
	var req struct {
		ID   uint
		Name string `binding:"max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var one admin.AdminRole
	one.ID = req.ID
	one.Name = req.Name

	if one.ID == 0 {
		one.AdminMenuApiPaths = "[]"
		one.AdminMenuIds = "[]"

		if err := bean.MysqlDb.Create(&one).Error; err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}
	} else {
		if err := bean.MysqlDb.Updates(&one).Error; err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}
	}

	c.String(200, "ok")
}

func RoleDelete(c *gin.Context) {
	// 查询列表
	var req struct {
		ID uint
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	// 判断是否在用
	var one admin.AdminRole
	if err := bean.MysqlDb.Where("id=?", req.ID).Take(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	if err := bean.MysqlDb.Delete(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.String(200, "ok")
}

// 修改权限
func RolePermissionEdit(c *gin.Context) {
	// 查询列表
	var req struct {
		ID           uint
		AdminMenuIds []uint
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	// 判断是否存在
	var one admin.AdminRole
	if err := bean.MysqlDb.Where("id=?", req.ID).Take(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(400, err.Error())
		return
	}

	// updates
	var one2 admin.AdminRole
	one2.ID = one.ID

	if menuIds, err := json.Marshal(req.AdminMenuIds); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	} else {
		one2.AdminMenuIds = string(menuIds)
	}

	// 查询 menu_id 转为 api_path
	var menu []admin.AdminMenu
	var paths []string

	if err := bean.MysqlDb.Where("id in (?)", req.AdminMenuIds).Where("`type`='接口'").Find(&menu).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	for _, v := range menu {
		paths = append(paths, v.Path)
	}

	if path, err := json.Marshal(paths); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	} else {
		one2.AdminMenuApiPaths = string(path)
	}

	// save
	if err := bean.MysqlDb.Updates(&one2).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.String(200, "ok")
}
