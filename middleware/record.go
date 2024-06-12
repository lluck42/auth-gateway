package middleware

import (
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model/admin"

	"github.com/gin-gonic/gin"
)

func Record(c *gin.Context) {

	c.Next()

	// 获取登录用户对象
	var adm admin.Admin
	if err := adm.Get(c); err != nil {
		bean.Log.Error(c.FullPath())
		bean.Log.Error(err)
		return
	}

	// 请求记录
	if c.Request.Method != "GET" {
		var rec admin.AdminRecord
		rec.AdminID = adm.ID
		rec.AdminName = adm.Name
		rec.Method = c.Request.Method
		rec.ApiPath = c.FullPath()
		rec.Ip = c.ClientIP()

		if err := bean.MysqlDb.Save(&rec).Error; err != nil {
			bean.Log.Error(err)
			return
		}
	}

}
