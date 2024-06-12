package admin

import (
	"encoding/json"
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model/admin"

	"github.com/gin-gonic/gin"
)

func MenuTree(c *gin.Context) {
	// 查询列表
	var req struct {
		Type string `binding:"oneof='' My 页面 接口"` // 菜单用
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var tx = bean.MysqlDb.Model(&admin.AdminMenu{})
	if req.Type == "My" {
		// 我的菜单
		tx.Where("`type`='页面' AND `show`='true'")
		// 获取登录用户对象
		var adm admin.Admin
		if err := adm.Get(c); err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}

		var rol admin.AdminRole
		if err := bean.MysqlDb.Where("id=?", adm.RoleID).Take(&rol).Error; err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}

		var menuIds []uint
		if err := json.Unmarshal([]byte(rol.AdminMenuIds), &menuIds); err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}

		// 待测试
		tx.Where("id in (?)", menuIds)
	} else if req.Type == "页面" {
		// 所有菜单
		tx.Where("`type`='页面'")
	}

	// 查询
	var list []admin.AdminMenu

	// 查询列表
	if err := tx.Order("`sort` ASC").Find(&list).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	var adm admin.AdminMenu

	c.PureJSON(200, adm.Tree(0, list))
}

func MenuDetail(c *gin.Context) {
	// 查询列表
	var req struct {
		ID int `binding:"required,gte=1"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var one admin.AdminMenu
	if err := bean.MysqlDb.First(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, one)
}

func MenuEdit(c *gin.Context) {
	// 查询列表
	var req struct {
		ID       uint
		ParentID uint
		Name     string `binding:"max=255"`
		Label    string `binding:"max=255"`
		Path     string `binding:"max=255"`
		Icon     string `binding:"max=255"`
		Type     string `binding:"oneof=页面 接口"`
		Show     string `binding:"oneof=true false"`
		Sort     uint
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var one admin.AdminMenu
	one.ID = req.ID
	one.ParentID = req.ParentID
	one.Name = req.Name
	one.Label = req.Label
	one.Path = req.Path
	one.Icon = req.Icon
	one.Type = req.Type
	one.Show = req.Show
	one.Sort = req.Sort

	if err := bean.MysqlDb.Save(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.String(200, "ok")
}

func MenuDelete(c *gin.Context) {
	// 查询列表
	var req struct {
		ID uint
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var one admin.AdminMenu
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
