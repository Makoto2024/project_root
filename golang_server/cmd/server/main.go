// Package main is the entry point for the server.
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/golang/glog"

	"github.com/Makoto2024/project_root/golang_server/internal/server"

	pb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

// cfg will be set via flags.
type cfg struct {
	IP   string
	Port int
}

func (c *cfg) registerFlags(fs *flag.FlagSet) {
	if fs == nil {
		fs = flag.CommandLine
	}
	fs.StringVar(&c.IP, "ip", "0.0.0.0", "IP address to listen on")
	fs.IntVar(&c.Port, "port", 8080, "Port to listen on")
}

func mainWithCfg(ctx context.Context, cfg *cfg) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.IP, cfg.Port))
	if err != nil {
		glog.ExitContextf(ctx, "Failed to listen: %v", err)
	}

	var serverOpts []grpc.ServerOption
	grpcServer := grpc.NewServer(serverOpts...)
	pb.RegisterHelloServiceServer(grpcServer, server.NewHelloServer())
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
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
