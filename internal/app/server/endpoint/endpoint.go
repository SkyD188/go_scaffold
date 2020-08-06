package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	"micro_service/api/pb"
)

// 业务逻辑
type MyService struct{}

func (svc MyService) Echo(ctx context.Context, req *pb.EchoRequest) (resp *pb.EchoResponse, err error) {
	return &pb.EchoResponse{Pong: "Echo: " + req.Ping}, nil
}

func MakeEchoEndpoint(svc pb.MyServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EchoRequest)
		return svc.Echo(ctx, req)
	}
}

// 一组endpoint
type EndpointSet struct {
	Echo endpoint.Endpoint
}
