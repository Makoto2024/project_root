package server

import (
	"context"

	"google.golang.org/protobuf/proto"

	pb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

var _ pb.PongServiceServer = (*PongServer)(nil) // Static check.

// PongServer implements the PongService.
type PongServer struct {
	pb.UnimplementedPongServiceServer
}

func (s *PongServer) Pong(ctx context.Context, req *pb.PongRequest) (*pb.PongResponse, error) {
	return pb.PongResponse_builder{
		Msg: proto.String("pong"),
	}.Build(), nil
}

// NewPongServer creates a new PongServer.
func NewPongServer() *PongServer {
	return &PongServer{}
}
