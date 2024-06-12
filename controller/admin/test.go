package admin

import (
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	c.PureJSON(200, 11)
}
