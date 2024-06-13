package middleware

import (
	"auth-gateway/bean"
	"auth-gateway/model/admin"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Permission(c *gin.Context) {
	fmt.Println("Permission")
	// 获取登录用户对象
	var adm admin.Admin
	if err := adm.Get(c); err != nil {
		bean.Log.Error(c.Request.URL.Path)
		bean.Log.Error(err)

		c.String(403, err.Error()+c.Request.URL.Path)
		c.Abort()
		return
	}

	fmt.Println("c.Request.URL.Path", c.Request.URL.Path)
	var role admin.AdminRole
	if err := bean.MysqlDb.Where("`id`=?", adm.RoleID).Where("JSON_OVERLAPS(`admin_menu_api_paths`, JSON_ARRAY(?))", c.Request.URL.Path).Take(&role).Error; err != nil {
		bean.Log.Error(err)
		c.String(403, "Permission 无权限 "+c.Request.URL.Path)
		c.Abort()
		return
	}
}
