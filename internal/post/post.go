package post

import (
	"context"
	"errors"
)

var (
	ErrEmptyTitle   = errors.New("title cannot be empty")
	ErrEmptyContent = errors.New("content cannot be empty")
)

type blogServer struct {
	UnimplementedBlogServer
}

func NewBlogServer() blogServer {
	return blogServer{}
}

func (s blogServer) CreatePost(ctx context.Context, request *CreatePostRequest) (*CreatePostResponse, error) {
	if request.Title == "" {
		return nil, ErrEmptyTitle
	}
	if request.Content == "" {
		return nil, ErrEmptyContent
	}
	return nil, nil
}
