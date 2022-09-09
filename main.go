package main

import (
	"echo/app/controllers/echo"
	pb "echo/grpc/proto"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"net"
)

type UserInfoService struct {
	pb.UnimplementedUserInfoServiceServer
}

var u = UserInfoService{}

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

	// 1. 监听
	address := "0.0.0.0:30005"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("监听异常 ", err)
		return
	} else {
		fmt.Println("监听成功:", address)
	}
	// 2. 实例化rpc
	s := grpc.NewServer()
	// 3. 在gRPC上注册微服务
	// 第一个类型是服务，第二个类型是接口的变量
	pb.RegisterUserInfoServiceServer(s, &u)
	// 4. 启动gRPC服务端
	_ = s.Serve(lis)
}
