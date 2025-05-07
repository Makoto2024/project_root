package server

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	ccpb "github.com/Makoto2024/project_root/cc_server/protos/servicepb"
	gopb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

var pongServiceCfg = struct {
	Target     string
	Credential credentials.TransportCredentials
}{
	Target:     "cc-server:7070",
	Credential: insecure.NewCredentials(),
}

var _ gopb.PongServiceServer = (*PongServer)(nil) // Static check.

// PongServer implements the PongService.
type PongServer struct {
	gopb.UnimplementedPongServiceServer
}

func (s *PongServer) Pong(ctx context.Context, req *gopb.PongRequest) (*gopb.PongResponse, error) {
	// Call PingService.Ping with `{req.Msg};pong` and return rsp.Msg.
	conn, err := newCCServerClient()
	if err != nil {
		return nil, fmt.Errorf("newCCServerClient(%q): %w", pongServiceCfg.Target, err)
	}
	defer conn.Close()

	client := ccpb.NewPingServiceClient(conn)
	reqStr := fmt.Sprintf("%s;pong", req.GetMsg())
	rsp, err := client.Ping(ctx, ccpb.PingRequest_builder{
		Msg: proto.String(reqStr),
	}.Build())
	if err != nil {
		return nil, fmt.Errorf("PongService.Pong(req.Msg = %q): %w", reqStr, err)
	}
	return gopb.PongResponse_builder{
		Msg: proto.String(rsp.GetMsg()),
	}.Build(), nil
}

// NewPongServer creates a new PongServer.
func NewPongServer() *PongServer {
	return &PongServer{}
}

func newCCServerClient() (*grpc.ClientConn, error) {
	return grpc.NewClient(
		pongServiceCfg.Target,
		grpc.WithTransportCredentials(pongServiceCfg.Credential),
	)
}
