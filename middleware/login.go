package middleware

import (
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model/admin"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var token = c.GetHeader("authorization")

	if len(token) < 7 {
		bean.Log.Error("token 错误l1")
		c.String(401, "token 错误l1")
		c.Abort()
		return
	}

	token = token[7:]

	var one admin.Admin
	if err := bean.MysqlDb.Where("token=?", token).First(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(401, "token 错误l2")
		c.Abort()
		return
	}

	c.Set("admin", one)
}
