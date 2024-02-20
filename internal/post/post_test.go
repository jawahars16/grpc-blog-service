package post_test

import (
	"context"
	"testing"
	"time"

	"github.com/jawahars16/grpc-blog-service/internal/post"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_CreatePost(t *testing.T) {
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
}
