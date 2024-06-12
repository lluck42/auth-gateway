package admin

import (
	"auth-gateway/bean"
	"auth-gateway/model/admin"

	"github.com/gin-gonic/gin"
)

func RecordList(c *gin.Context) {

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
	var list []admin.AdminRecord

	var tx = bean.MysqlDb.Model(&admin.AdminRecord{})

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
