package post

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrEmptyTitle   = status.Error(codes.InvalidArgument, "title cannot be empty")
	ErrEmptyContent = status.Error(codes.InvalidArgument, "content cannot be empty")
	ErrCreatePost   = status.Error(codes.Internal, "failed to create post")
	ErrUpdatePost   = status.Error(codes.Internal, "failed to update post")
	ErrEmptyPostID  = status.Error(codes.InvalidArgument, "post id cannot be empty")
	ErrPostNotFound = status.Error(codes.NotFound, "post not found")
)

type storage interface {
	Set(id string, item interface{}) error
	Get(id string) (interface{}, bool)
	Delete(id string) bool
}

type blogServer struct {
	UnimplementedBlogServer
	data storage
}

func NewBlogServer(storage storage) blogServer {
	return blogServer{
		data: storage,
	}
}

func (s blogServer) CreatePost(ctx context.Context, request *CreatePostRequest) (*CreatePostResponse, error) {
	if request.Title == "" {
		return nil, ErrEmptyTitle
	}
	if request.Content == "" {
		return nil, ErrEmptyContent
	}
	id := uuid.New().String()
	post := &Post{
		PostId:          id,
		Title:           request.Title,
		Content:         request.Content,
		Author:          request.Author,
		PublicationDate: request.PublicationDate,
		Tags:            request.Tags,
	}
	err := s.data.Set(id, post)
	if err != nil {
		return nil, ErrCreatePost
	}
	return &CreatePostResponse{
		Post: post,
	}, nil
}

func (s blogServer) GetPost(ctx context.Context, request *GetPostRequest) (*GetPostResponse, error) {
	if request.PostId == "" {
		return nil, ErrEmptyPostID
	}
	postData, found := s.data.Get(request.PostId)
	if !found {
		return nil, ErrPostNotFound
	}
	post := postData.(*Post)
	return &GetPostResponse{
		Post: &Post{
			PostId:          post.PostId,
			Title:           post.Title,
			Content:         post.Content,
			Author:          post.Author,
			PublicationDate: post.PublicationDate,
			Tags:            post.Tags,
		},
	}, nil
}

func (s blogServer) UpdatePost(ctx context.Context, request *UpdatePostRequest) (*UpdatePostResponse, error) {
	if request.PostId == "" {
		return nil, ErrEmptyPostID
	}
	existingPostItem, found := s.data.Get(request.PostId)
	if !found {
		return nil, ErrPostNotFound
	}
	existingPost := existingPostItem.(*Post)
	updatedPost := &Post{
		PostId:          existingPost.PostId,
		Title:           request.Title,
		Content:         request.Content,
		Author:          request.Author,
		Tags:            request.Tags,
		PublicationDate: existingPost.PublicationDate,
	}
	err := s.data.Set(request.PostId, updatedPost)
	if err != nil {
		return nil, ErrUpdatePost
	}
	return &UpdatePostResponse{
		Post: updatedPost,
	}, nil
}

func (s blogServer) DeletePost(ctx context.Context, request *DeletePostRequest) (*DeletePostResponse, error) {
	if request.PostId == "" {
		return nil, ErrEmptyPostID
	}
	found := s.data.Delete(request.PostId)
	if !found {
		return nil, ErrPostNotFound
	}
	return &DeletePostResponse{}, nil
}
