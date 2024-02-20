package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/jawahars16/grpc-blog-service/internal/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	port := flag.Int("port", 9000, "Port to connect to")
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Error in creating connection.", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	client := post.NewBlogClient(conn)
	res, err := client.CreatePost(context.Background(), &post.CreatePostRequest{
		Title:   "Test Title",
		Content: "Test Content",
	})
	if err != nil {
		fmt.Println("Failed to create post.", err)
		return
	}
	fmt.Println("Post created successfully.", res.Post)
}
