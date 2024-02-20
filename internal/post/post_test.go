package post_test

import (
	"context"
	"testing"
	"time"

	"github.com/jawahars16/grpc-blog-service/internal/post"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Blog_Server(t *testing.T) {
	t.Run("given a request with an empty title, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer()
		_, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title: "",
		})
		assert.Error(t, err)
	})

	t.Run("given a request with an empty content, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer()
		_, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:   "Test Title",
			Content: "",
		})
		assert.Error(t, err)
	})

	t.Run("given a valid request, it should return a post with the same title and content", func(t *testing.T) {
		t.Skip("WIP")
		ctx := context.Background()
		server := post.NewBlogServer()
		response, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:           "Test Title",
			Content:         "Test Content",
			Author:          "Test Author",
			PublicationDate: timestamppb.New(time.Now()),
			Tags:            []string{"tag1", "tag2"},
		})

		assert.NoError(t, err)
		if assert.NotNil(t, response.Post) {
			assert.Equal(t, "Test Title", response.Post.Title)
			assert.Equal(t, "Test Content", response.Post.Content)
		}
	})
}
