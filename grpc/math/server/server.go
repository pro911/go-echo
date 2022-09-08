package server

import (
	"context"
	"echo/grpc/math/pd"
)

type MathService struct{}

func (m *MathService) Sub(ctx context.Context, request *pd.SubRequest) (*pd.SubResponse, error) {
	//TODO implement me
	var sum int32 = 0
	for k, arg := range request.Values {
		if k == 0 {
			sum = arg
		} else {
			sum = sum - arg
		}
	}
	return &pd.SubResponse{
		Val: sum,
	}, nil
}

func (m *MathService) Add(ctx context.Context, request *pd.AddRequest) (*pd.AddResponse, error) {
	var sum int32 = 0
	for _, arg := range request.Values {
		sum = sum + arg
	}
	return &pd.AddResponse{
		Val: sum,
	}, nil
}

//func main() {
//	server := grpc.NewServer()
//	pd.RegisterMathServiceServer(server, &MathService{})
//	lis, err := net.Listen("tcp", ":18083")
//	if err != nil {
//		panic(err.Error())
//	}
//	if err = server.Serve(lis); err != nil {
//		panic(err.Error())
//	}
//}
