package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

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

	loop(context.Background(), post.NewBlogClient(conn))
}

func loop(ctx context.Context, client post.BlogClient) {
	for {
		var input string
		fmt.Println("")
		fmt.Println("Enter command to interact with blogging platform. Type 'exit' to quit.")
		fmt.Println("[Commands: create, get, list, delete, exit]")
		fmt.Print("> ")
		fmt.Scanln(&input)

		switch input {
		case "exit":
			fmt.Println("Exiting...")
			return
		case "create":
			postDetails := readPostDetails()
			response, err := client.CreatePost(ctx, &post.CreatePostRequest{
				Title:   postDetails.Title,
				Content: postDetails.Content,
				Author:  postDetails.Author,
				Tags:    postDetails.Tags,
			})
			if err != nil {
				fmt.Println("Failed to create post.", err)
				continue
			}
			fmt.Println("Post created successfully.", response.Post)
			continue
		case "get":
			var input string
			fmt.Print("Post ID: ")
			fmt.Scanln(&input)
			response, err := client.GetPost(ctx, &post.GetPostRequest{
				PostId: input,
			})
			if err != nil {
				fmt.Println("Error fetching post.", err)
			}
			fmt.Println(response.Post)
			continue
		default:
			fmt.Println("Invalid input.")
			continue
		}
	}
}

func readPostDetails() *post.Post {
	scanner := bufio.NewScanner(os.Stdin)
	post := &post.Post{}
	fmt.Print("Title: ")
	if scanner.Scan() {
		post.Title = scanner.Text()
	}

	fmt.Print("Content: ")
	if scanner.Scan() {
		post.Content = scanner.Text()
	}

	fmt.Print("Author: ")
	if scanner.Scan() {
		post.Author = scanner.Text()
	}

	fmt.Print("Tags (comma separated): ")
	if scanner.Scan() {
		post.Tags = strings.Split(scanner.Text(), ",")
	}

	fmt.Println("")
	return post
}
