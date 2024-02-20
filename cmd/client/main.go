package main

import (
	"context"
	"fmt"

	"github.com/jawahars16/grpc-blog-service/internal/post"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	port := 9000
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("Failed to connect to server\n", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	client := post.NewBlogClient(conn)
	res, err := client.CreatePost(ctx, &post.CreatePostRequest{
		Title: "Test Title",
	})
	if err != nil {
		fmt.Println("Failed to create post\n", err)
		return
	}
	fmt.Println("Post created successfully\n", res.Post)
}
