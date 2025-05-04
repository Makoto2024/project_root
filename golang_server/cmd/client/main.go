// Package main is is the entry point for the client.
package main

import (
	"context"
	"flag"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/golang/glog"

	pb "github.com/Makoto2024/project_root/golang_server/protos/servicepb"
)

// cfg will be set via flags.
type cfg struct {
	IP   string
	Port int
	Name string
}

func (c *cfg) registerFlags(fs *flag.FlagSet) {
	if fs == nil {
		fs = flag.CommandLine
	}
	fs.StringVar(&c.IP, "ip", "127.0.0.1", "IP address to connect to")
	fs.IntVar(&c.Port, "port", 8080, "Port to connect to")
	fs.StringVar(&c.Name, "name", "gRPC", "Name to use in the request")
}

func mainWithCfg(ctx context.Context, cfg *cfg) {
	target := fmt.Sprintf("%s:%d", cfg.IP, cfg.Port)
	credOpt := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient(target, credOpt)
	if err != nil {
		glog.FatalContextf(ctx, "grpc.Dial(%q) err: %v", target, err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	resp, err := client.Hello(ctx, pb.HelloRequest_builder{Name: &cfg.Name}.Build())
	if err != nil {
		glog.FatalContextf(ctx, "client.Search err: %v", err)
	}

	glog.ErrorContextf(ctx, "resp.greeting: %s", resp.GetGreeting())
}

func main() {
	var cfg cfg
	cfg.registerFlags(nil)
	flag.Parse()

	ctx := context.Background()
	glog.ErrorContextf(ctx, "cfg = %#v", cfg)
	mainWithCfg(ctx, &cfg)
}
