package config

import (
	"gq_volunteer_platform/bean"
	"gq_volunteer_platform/model/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Detail(c *gin.Context) {

	var id = 1

	var one config.Config
	if err := bean.MysqlDb.Where("id=?", id).Take(&one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, "数据库错误:id"+strconv.Itoa(id))
		return
	}

	c.PureJSON(200, one)
}

func Edit(c *gin.Context) {

	var req struct {
		Project string
		Mobile  string
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		bean.Log.Warning(err)
		c.String(400, err.Error())
		return
	}

	var one config.Config
	one.ID = 1
	one.Project = req.Project
	one.Mobile = req.Mobile

	//
	if err := bean.MysqlDb.Updates(one).Error; err != nil {
		bean.Log.Error(err)
		c.String(500, "数据库错误:id"+strconv.Itoa(int(one.ID)))
		return
	}

	c.String(200, "ok")
}
