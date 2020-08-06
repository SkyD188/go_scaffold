package main

import (
	"flag"
	"net"
	"os"
	"time"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"

	"micro_service/api/pb"
	"micro_service/internal/app/server/endpoint"
	"micro_service/internal/app/server/transport"
)

func main() {
	var (
		grpcAddr            = flag.String("grpc-addr", ":8081", "gRPC listen address")
		logger              = log.NewLogfmtLogger(os.Stderr)
		loggingMiddleware   = endpoint.LoggingMiddleware(logger)
		limiter             = rate.NewLimiter(rate.Every(time.Second), 100)
		rateLimitMiddleware = endpoint.RateLimitMiddleware(limiter)
	)
	flag.Parse()

	service := endpoint.MyService{}
	endpointSet := endpoint.EndpointSet{Echo: endpoint.MakeEchoEndpoint(service)}

	grpcServer := transport.NewGrpcServer(endpointSet, loggingMiddleware, rateLimitMiddleware)
	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
	pb.RegisterMyServiceServer(baseServer, grpcServer)

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		logger.Log(err)
	}

	logger.Log(baseServer.Serve(grpcListener))
}
