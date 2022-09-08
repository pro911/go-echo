package server

import (
	"context"
	"echo/grpc/send/pd"
)

type SendService struct {
	*pd.UnimplementedSendServiceServer
}

func (m *SendService) Send(ctx context.Context, request *pd.Request) (*pd.Response, error) {
	return &pd.Response{
		MsgCode: request.MsgCode,
		MsgData: request.MsgData,
	}, nil
}
