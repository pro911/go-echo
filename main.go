package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	serve := gin.Default()
	serve.Any("/", indexHandler)
	serve.Any("/baogao", func(c *gin.Context) {
		c.File("baogao.jpg")
	})
	if err := serve.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n\n", err)
	}
}

func indexHandler(c *gin.Context) {

	//file, _ := c.FormFile("file")

	var header = make(map[string]string)
	for k, v := range c.Request.Header {
		var st string
		for i, s := range v {
			if i == 0 {
				st += s
			} else {
				st += ";" + s
			}
		}
		header[k] = st
	}

	body, _ := c.GetRawData()

	username, pw, _ := c.Request.BasicAuth()

	c.JSON(http.StatusOK, gin.H{
		"header": header,
		"query":  c.Request.URL.RawQuery,
		"body":   string(body),
		"url":    c.Request.URL,
		"method": c.Request.Method,
		"proto":  c.Request.Proto,
		"tls":    c.Request.TLS,
		"basicAuth": map[string]string{
			"username": username,
			"password": pw,
		},
	})
}
