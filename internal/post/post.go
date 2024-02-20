package post

type blogServer struct {
}

func NewBlogServer() blogServer {
	return blogServer{}
}

func (s blogServer) CreatePost(request *CreatePostRequest) (*CreatePostResponse, error) {
	return &CreatePostResponse{
		Post: &Post{
			Title:   request.Title,
			Content: request.Content,
		},
	}, nil
}
