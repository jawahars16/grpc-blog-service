package post_test

import (
	"context"
	"testing"
	"time"

	"github.com/jawahars16/grpc-blog-service/internal/data"
	"github.com/jawahars16/grpc-blog-service/internal/post"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Test_Blog_Server_CreatePost(t *testing.T) {
	t.Run("given a request with an empty title, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title: "",
		})
		assert.Error(t, err)
	})

	t.Run("given a request with an empty content, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:   "Test Title",
			Content: "",
		})
		assert.Error(t, err)
	})

	t.Run("given a valid request, it should return a post with the same title and content", func(t *testing.T) {
		ctx := context.Background()
		publicationTime := timestamppb.New(time.Now())
		server := post.NewBlogServer(data.NewInMemoryStorage())
		response, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:           "Test Title",
			Content:         "Test Content",
			Author:          "Test Author",
			PublicationDate: publicationTime,
			Tags:            []string{"tag1", "tag2"},
		})

		assert.NoError(t, err)
		if assert.NotNil(t, response.Post) {
			assert.Equal(t, "Test Title", response.Post.Title)
			assert.Equal(t, "Test Content", response.Post.Content)
			assert.Equal(t, "Test Author", response.Post.Author)
			assert.Equal(t, publicationTime, response.Post.PublicationDate)
			assert.Equal(t, []string{"tag1", "tag2"}, response.Post.Tags)
			assert.NotEmpty(t, response.Post.PostId)
		}
	})
}

func Test_Blog_Server_GetPost(t *testing.T) {
	t.Run("given a request with an empty post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.GetPost(ctx, &post.GetPostRequest{
			PostId: "",
		})
		assert.ErrorIs(t, err, post.ErrEmptyPostID)
	})

	t.Run("given a request with a non-existing post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.GetPost(ctx, &post.GetPostRequest{
			PostId: "non-existing-id",
		})
		assert.ErrorIs(t, err, post.ErrPostNotFound)
	})

	t.Run("given a request with an existing post id, it should return the post", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		publicationTime := timestamppb.New(time.Now())
		createResponse, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:           "Test Title",
			Content:         "Test Content",
			PublicationDate: publicationTime,
		})
		assert.NoError(t, err)
		getResponse, err := server.GetPost(ctx, &post.GetPostRequest{
			PostId: createResponse.Post.PostId,
		})
		assert.NoError(t, err)
		if assert.NotNil(t, getResponse.Post) {
			assert.Equal(t, "Test Title", getResponse.Post.Title)
			assert.Equal(t, "Test Content", getResponse.Post.Content)
			assert.Equal(t, publicationTime, getResponse.Post.PublicationDate)
		}
	})
}

func Test_Blog_Server_UpdatePost(t *testing.T) {
	t.Run("given a request with an empty post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.UpdatePost(ctx, &post.UpdatePostRequest{
			PostId: "",
			Title:  "Test title",
		})
		assert.ErrorIs(t, err, post.ErrEmptyPostID)
	})

	t.Run("given a request with a non-existing post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.UpdatePost(ctx, &post.UpdatePostRequest{
			PostId: "non-existent-id",
			Title:  "Test title",
		})
		assert.ErrorIs(t, err, post.ErrPostNotFound)
	})

	t.Run("given a request with an existing post id, it should return the updated post", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		createResponse, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:   "Title",
			Content: "Content",
		})
		assert.NoError(t, err)
		_, err = server.UpdatePost(ctx, &post.UpdatePostRequest{
			PostId: createResponse.Post.PostId,
			Title:  "Modified title",
		})
		assert.NoError(t, err)
		getResponse, err := server.GetPost(ctx, &post.GetPostRequest{
			PostId: createResponse.Post.PostId,
		})
		assert.NoError(t, err)
		if assert.NotNil(t, getResponse.Post) {
			assert.Equal(t, "Modified title", getResponse.Post.Title)
		}
	})
}

func Test_Blog_Server_DeletePost(t *testing.T) {
	t.Run("given a request with an empty post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.DeletePost(ctx, &post.DeletePostRequest{
			PostId: "",
		})
		assert.ErrorIs(t, err, post.ErrEmptyPostID)
	})

	t.Run("given a request with a non-existing post id, it should return an error", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		_, err := server.DeletePost(ctx, &post.DeletePostRequest{
			PostId: "non-existent-id",
		})
		assert.ErrorIs(t, err, post.ErrPostNotFound)
	})

	t.Run("given a request with an existing post id, it should delete the post", func(t *testing.T) {
		ctx := context.Background()
		server := post.NewBlogServer(data.NewInMemoryStorage())
		createResponse, err := server.CreatePost(ctx, &post.CreatePostRequest{
			Title:   "Title",
			Content: "Content",
		})
		assert.NoError(t, err)
		_, err = server.DeletePost(ctx, &post.DeletePostRequest{
			PostId: createResponse.Post.PostId,
		})
		assert.NoError(t, err)
		_, err = server.GetPost(ctx, &post.GetPostRequest{
			PostId: createResponse.Post.PostId,
		})
		assert.ErrorIs(t, err, post.ErrPostNotFound)
	})
}
