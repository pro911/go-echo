package main

import (
	"echo/app/controllers/echo"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	serve := gin.Default()
	serve.Any("/", echo.IndexHandler)
	serve.Any("/json", echo.JsonDecode)
	serve.Any("/baogao", func(c *gin.Context) {
		c.File("baogao.jpg")
	})
	if err := serve.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}
