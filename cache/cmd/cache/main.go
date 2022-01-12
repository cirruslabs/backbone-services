package main

import (
	"flag"
	"fmt"
	"github.com/cirruslabs/backbone-services/cache/internal/impl"
	"github.com/cirruslabs/backbone-services/proto/cache"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	cache.RegisterCacheServer(grpcServer, impl.NewInmemory())
	grpcServer.Serve(lis)
}
