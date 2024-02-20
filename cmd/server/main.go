package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/jawahars16/grpc-blog-service/internal/data"
	"github.com/jawahars16/grpc-blog-service/internal/post"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 9000, "Port to connect to")
	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal("Sever failed to start\n", err)
	}

	grpcServer := grpc.NewServer()
	post.RegisterBlogServer(grpcServer, post.NewBlogServer(data.NewInMemoryStorage()))

	log.Printf("Server started on port %d\n", *port)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("gRPC server failed to start\n", err)
	}
}
