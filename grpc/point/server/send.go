package server

import (
	"context"
	"echo/grpc/point/pd"
)

type AppPointService struct {
	*pd.UnimplementedAppPointServiceServer
}

func (m *AppPointService) Send(ctx context.Context, request *pd.Request) (*pd.Response, error) {
	return &pd.Response{
		MsgCode: request.MsgCode,
		MsgData: request.MsgData,
	}, nil
}
