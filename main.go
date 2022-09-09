package main

import (
	"echo/app/controllers/echo"
	pb "echo/grpc/proto"
	"echo/grpc/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strings"
)

var (
	g errgroup.Group
)

var u = service.UserInfoService{}

func main() {

	grpcServe := grpc.NewServer()

	pb.RegisterUserInfoServiceServer(grpcServe, &u)

	httpServe := gin.New()

	httpServe.Any("/", echo.IndexHandler)
	httpServe.GET("/websocket", echo.Websocket)
	httpServe.Any("/json", echo.JsonDecode)
	httpServe.Any("/img_base", echo.ImgBase)
	httpServe.Any("/status", echo.ReturnStatus)
	httpServe.Any("/baogao", func(c *gin.Context) {
		c.File("baogao.jpg")
	})

	// 监听端口并处理服务分流
	h2Handler := h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 判断协议是否为http/2 && 是grpc
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
			//按照grpc处理请求
			grpcServe.ServeHTTP(w, r)
		} else {
			//当作普通api处理
			httpServe.ServeHTTP(w, r)
		}
	}), &http2.Server{})

	// 监听HTTP服务
	if err := http.ListenAndServe(":8000", h2Handler); err != nil {
		log.Println("http server done:", err.Error())
	}
}
