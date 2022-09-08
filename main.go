package main

import (
	"echo/app/controllers/echo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	serve := gin.Default()
	serve.Any("/", echo.IndexHandler)
	serve.GET("/websocket", echo.Websocket)
	serve.Any("/json", echo.JsonDecode)
	serve.Any("/img_base", echo.ImgBase)
	serve.Any("/status", echo.ReturnStatus)
	serve.Any("/baogao", func(c *gin.Context) {
		c.File("baogao.jpg")
	})
	if err := serve.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
