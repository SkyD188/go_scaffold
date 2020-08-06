package myendpoint

import (
	"context"
	"log"
	"os"

	"google.golang.org/grpc"

	"micro_service/api/pb"
)

var address string = "10.221.113.184:8081"

// 调用rpc
func RpcResp() string {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyServiceClient(conn)

	if len(os.Args) > 1 {
		log.Println(os.Args[1])
	}
	r, err := c.Echo(context.Background(), &pb.EchoRequest{Ping: "ping"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return r.Pong
}
