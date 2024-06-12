package admin

import (
	"auth-gateway/bean"
	"auth-gateway/model/admin"

	"github.com/gin-gonic/gin"
)

func MyList(c *gin.Context) {

	var req struct {
		Page           int    `binding:"required,gte=1"`
		Limit          int    `binding:"required,gte=1"`
		Name           string `binding:"max=255"`
		OrganizationId uint   `binding:"max=10000000"`
	}

	if err := c.ShouldBindQuery(&req); err != nil {
		bean.Log.Warning(err.Error())
		c.String(400, err.Error())
		return
	}

	var offset = (req.Page - 1) * req.Limit

	// 查询
	var count int64
	var list []admin.Admin

	var tx = bean.MysqlDb.Model(&admin.Admin{})

	if req.Name != "" {
		tx = tx.Where("`name` like ?", "%"+req.Name+"%")
	}

	// 查询总数
	if err := tx.Count(&count).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	// 查询列表
	if err := tx.Limit(req.Limit).Offset(offset).Order("`id` DESC").Find(&list).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, gin.H{
		"count": count,
		"list":  list,
	})
}

// info
func Info(c *gin.Context) {
	// 获取登录用户对象
	var one admin.Admin
	if err := one.Get(c); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	c.PureJSON(200, one)
}

// 添加+修改管理员信息
func Edit(c *gin.Context) {
	// 用户输入信息
	var req struct {
		ID             uint   `binding:"max=10000000"`
		Name           string `binding:"max=255"`
		Mobile         string `binding:"max=255"`
		Role           string `binding:"max=255"`
		RoleID         uint   `binding:"min=1,max=10000000"`
		OrganizationId uint   `binding:"min=1,max=10000000"`
		Password       string `binding:"max=255"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err)
		c.String(400, err.Error())
		return
	}

	if req.ID == 0 && req.Password == "" {
		bean.Log.Error("密码不能为空")
		c.String(400, "密码不能为空")
		return
	}

	var one admin.Admin
	one.ID = req.ID
	one.Name = req.Name
	one.Mobile = req.Mobile
	one.RoleID = req.RoleID
	one.Role = req.Role
	one.OrganizationId = req.OrganizationId
	one.Password = req.Password

	if one.Password != "" {
		if err := one.PasswordMd5(); err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}

		// 重置token
		if err := one.NewToken(); err != nil {
			bean.Log.Error(err)
			c.String(500, err.Error())
			return
		}
	}

	if req.ID == 0 {
		if err := bean.MysqlDb.Save(&one).Error; err != nil {
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

// 删除管理员
func Delete(c *gin.Context) {
	// 用户输入信息
	var req struct {
		ID int
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err)
		c.String(400, err.Error())
		return
	}

	if req.ID == 1 {
		c.String(400, "管理员不能删除")
		return
	}

	var one admin.Admin
	if err := bean.MysqlDb.Where("id=?", req.ID).First(&one).Error; err != nil {
		bean.Log.Warning(err)
		c.String(400, "没有该用户")
		return
	}

	if err := bean.MysqlDb.Delete(&one).Error; err != nil {
		bean.Log.Warning(err)
		c.String(500, "内部错误")
		return
	}

	c.String(200, "ok")
}

func ChangePassword(c *gin.Context) {
	// 用户输入信息
	var req struct {
		Password    string
		NewPassword string
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Error(err)
		c.String(400, err.Error())
		return
	}

	// 获取登录用户对象
	var one admin.Admin
	if err := one.Get(c); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	}

	// 新旧密码不能相同
	if req.Password == req.NewPassword {
		c.String(400, "新旧密码相同")
		return
	}

	// 验证旧密码
	if s, err := admin.Md5(req.Password); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	} else {
		if s != one.Password {
			c.String(400, "密码错误")
			return
		}
	}

	// 加密新密码
	var str = req.NewPassword
	if s, err := admin.Md5(str); err != nil {
		bean.Log.Error(err)
		c.String(500, err.Error())
		return
	} else {
		one.Password = s
	}

	if err := bean.MysqlDb.Save(&one).Error; err != nil {
		c.String(500, err.Error())
		return
	}

	c.String(200, "ok")
}
