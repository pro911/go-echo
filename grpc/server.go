package main

import (
	"context"
	pb "echo/grpc/proto"
	"errors"
)

type UserInfoService struct {
	pb.UnimplementedUserInfoServiceServer
}

var u = UserInfoService{}

// GetUserInfo 实现服务端需要实现的端口
func (service *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	name := req.Name
	// 在数据库查询用户信息
	if name != "" {
		err = nil
	} else {
		err = errors.New("name 不能为空")
	}
	resp = &pb.UserResponse{
		Id:   1,
		Name: name,
		Age:  19,
		// 切片字段
		Hobby: []string{"FanOne", "FanOneTwo"},
	}
	return
}

func (service *UserInfoService) GetUserName(ctx context.Context, req *pb.NameRequest) (resp *pb.NameResponse, err error) {
	id := req.Id
	// 在数据库查询用户信息
	if id == 1 {
		resp = &pb.NameResponse{
			Name: "apipost.cn",
		}
	} else {
		resp = &pb.NameResponse{
			Name: "apipost.cn",
		}
	}

	err = nil
	return
}
