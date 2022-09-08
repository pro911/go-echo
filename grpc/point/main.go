package main

import (
	"echo/grpc/point/pd"
	server2 "echo/grpc/point/server"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	port := ":58083"
	for i, v := range os.Args {
		if i == 1 {
			port = v
		}
	}
	server := grpc.NewServer()
	pd.RegisterAppPointServiceServer(server, &server2.AppPointService{})
	fmt.Println("grpc port", port)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err.Error())
	}
	if err = server.Serve(lis); err != nil {
		panic(err.Error())
	}
}
