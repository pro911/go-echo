package main

import (
	"echo/app/controllers/echo"
	pb "echo/grpc/proto"
	"echo/grpc/service"
	resource "echo/resource/example"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"strings"

	//"io/ioutil"
	"log"
	"net/http"
)

//调用具体的grpc service
var u = service.UserInfoService{}

func main() {

	//1.NewServer 创建一个未注册服务且尚未开始接受请求的 gRPC 服务器
	grpcServe := grpc.NewServer()

	//2.注册具体的grpc服务
	pb.RegisterUserInfoServiceServer(grpcServe, &u)

	//3.New 返回一个新的空白 Engine 实例，没有附加任何中间件。
	httpServe := gin.New()

	httpServe.Any("/", echo.IndexHandler)
	httpServe.GET("/websocket", echo.Websocket)
	httpServe.Any("/json", echo.JsonDecode)
	httpServe.Any("/img_base", echo.ImgBase)
	httpServe.Any("/status", echo.ReturnStatus)
	httpServe.Any("/baogao", func(c *gin.Context) {
		c.File("baogao.jpg")
	})

	httpServe.Any("/example", func(c *gin.Context) {
		//'terminal'
		terminal := c.GetHeader("terminal") //web || client

		if terminal == "" {
			terminal = "web"
		}

		//获取文件地址
		filePath := fmt.Sprintf("%v%v.json", "resources/static/example/", terminal)
		fmt.Println(filePath)
		//b, err := ioutil.ReadFile(filePath) //读取文件内容
		b, err := resource.Asset(filePath)
		if err != nil {
			c.JSON(http.StatusOK, err.Error())
		} else {
			c.String(http.StatusOK, string(b))
		}
		return
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
