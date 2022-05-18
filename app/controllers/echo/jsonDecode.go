package echo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TJsonDecode struct {
	Param interface{} `form:"param" json:"param" uri:"param" xml:"param"`
}

func JsonDecode(c *gin.Context) {

	// 声明接收的变量
	var request TJsonDecode
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&request); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		body, _ := c.GetRawData()
		c.JSON(http.StatusBadRequest, "参数错误"+string(body))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": request,
	})
}
