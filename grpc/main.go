package main

import (
	send "echo/grpc/send/pd"
	server2 "echo/grpc/send/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	server := grpc.NewServer()
	send.RegisterSendServiceServer(server, &server2.SendService{})
	send.RegisterSendServiceServer(server, &server2.SendService{})
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		panic(err.Error())
	}
	if err = server.Serve(lis); err != nil {
		panic(err.Error())
	}
}
