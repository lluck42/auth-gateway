package middleware

import (
	"auth-gateway/bean"
	"auth-gateway/model/admin"
	"strconv"

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

	c.Request.Header.Add("auth-gateway-adminID", strconv.Itoa(int(one.ID)))
}
