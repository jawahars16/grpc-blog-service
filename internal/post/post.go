package post

import (
	"context"
)

type blogServer struct {
	UnimplementedBlogServer
}

func NewBlogServer() blogServer {
	return blogServer{}
}

func (s blogServer) CreatePost(ctx context.Context, request *CreatePostRequest) (*CreatePostResponse, error) {
	return &CreatePostResponse{
		Post: &Post{
			Title:   request.Title,
			Content: request.Content,
		},
	}, nil
}
