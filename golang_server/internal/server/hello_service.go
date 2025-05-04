// Package server implements the golang_server's GRPC handlers.
package server

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	pb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

var _ pb.HelloServiceServer = (*HelloServer)(nil) // Static check.

// HelloServer implements the HelloService.
type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return pb.HelloResponse_builder{
		Greeting: proto.String(fmt.Sprintf("Hello %s!", req.GetName())),
	}.Build(), nil
}

// NewHelloServer creates a new HelloServer.
func NewHelloServer() *HelloServer {
	return &HelloServer{}
}
