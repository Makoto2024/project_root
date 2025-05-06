// Package main is is the entry point for the client.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/golang/glog"

	pb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

func callHello(
	ctx context.Context, cfg *cfg, conn *grpc.ClientConn,
) (*pb.HelloResponse, error) {
	client := pb.NewHelloServiceClient(conn)
	return client.Hello(ctx, pb.HelloRequest_builder{Name: &cfg.Param}.Build())
}

func callPong(
	ctx context.Context, cfg *cfg, conn *grpc.ClientConn,
) (*pb.PongResponse, error) {
	client := pb.NewPongServiceClient(conn)
	return client.Pong(ctx, pb.PongRequest_builder{Msg: &cfg.Param}.Build())
}

func mainWithCfg(ctx context.Context, cfg *cfg) {
	target := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)
	credOpt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(target, credOpt)
	if err != nil {
		glog.ExitContextf(ctx, "grpc.Dial(%q) err: %v", target, err)
	}
	defer conn.Close()

	switch cfg.Service {
	case serviceTypeHelloService:
		resp, err := callHello(ctx, cfg, conn)
		if err != nil {
			glog.ExitContextf(ctx, "HelloService.Hello = %v", err)
		}
		glog.ErrorContextf(ctx, "HelloService.Hello = %s", resp.GetGreeting())
	case serviceTypePongService:
		resp, err := callPong(ctx, cfg, conn)
		if err != nil {
			glog.ExitContextf(ctx, "PongService.Pong = %v", err)
		}
		glog.ErrorContextf(ctx, "PongService.Pong = %s", resp.GetMsg())
	}
}

func main() {
	ctx := context.Background()

	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	var cfg cfg
	cfg.registerFlags(fs)
	if err := fs.Parse(os.Args[1:]); err != nil {
		glog.ExitContextf(ctx, "Parse command line args = %v", err)
	}

	glog.ErrorContextf(ctx, "cfg = %#v", cfg)
	mainWithCfg(ctx, &cfg)
}
