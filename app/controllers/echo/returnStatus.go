package echo

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ReturnStatus(c *gin.Context) {
	s := c.GetHeader("code")
	if i, e := strconv.Atoi(s); e == nil {
		c.JSON(i, gin.H{})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"data": "参数code不能为空"})
}
