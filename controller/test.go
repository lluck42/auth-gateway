package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Time(c *gin.Context) {
	c.String(200, time.Now().Format(time.DateTime))
}
