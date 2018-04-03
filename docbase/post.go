package docbase

import (
	"context"
	"fmt"
)

// PostService handles communication with the Post related
// methods of the Docbase API.
type PostService service

// PostListOptions specifies the optional parameters to the
// PostService.List method.
type PostListOptions struct {
	// Query filters Post.
	//
	// docs: https://help.docbase.io/posts/59432?list=%2Fsearch&q=%E6%A4%9C%E7%B4%A2#%E6%A4%9C%E7%B4%A2%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3
	Query string `url:"q,omitempty"`

	ListOptions
}

// TODO: QueryBuilder

// List posts in a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92984
func (s *PostService) List(ctx context.Context, domain string, opt *PostListOptions) ([]Post, *Response, error) {
	u := fmt.Sprintf("teams/%v/posts", domain)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var posts struct {
		Meta  Meta   `json:"meta"`
		Posts []Post `json:"posts"`
	}
	resp, err := s.client.Do(ctx, req, &posts)
	if err != nil {
		return nil, resp, err
	}
	resp.Meta = posts.Meta

	return posts.Posts, resp, nil
}

// Get a single post.
//
// Docbase API docs: https://help.docbase.io/posts/97204
func (s *PostService) Get(ctx context.Context, domain string, id int64) (*Post, *Response, error) {
	u := fmt.Sprintf("teams/%v/posts/%v", domain, id)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	post := new(Post)
	resp, err := s.client.Do(ctx, req, post)
	if err != nil {
		return nil, resp, err
	}

	return post, resp, nil
}

// CreatePostOptions represents a Docbase Post.
type CreatePostOptions struct {
	Body   string   `json:"body"`
	Draft  bool     `json:"draft"`
	Groups []Group  `json:"groups"`
	Scope  string   `json:"scope"`
	Tags   []string `json:"tags"`
	Title  string   `json:"title"`
}

// Create a post for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92980
func (s *PostService) Create(ctx context.Context, domain string, post *CreatePostOptions) (*Post, *Response, error) {
	u := fmt.Sprintf("teams/%v/posts", domain)
	req, err := s.client.NewRequest("POST", u, post)
	if err != nil {
		return nil, nil, err
	}

	g := new(Post)
	resp, err := s.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Edit a post.
//
// Docbase API docs: https://help.docbase.io/posts/92981
func (s *PostService) Edit(ctx context.Context, domain string, id int64, post *Post) (*Post, *Response, error) {
	u := fmt.Sprintf("teams/%v/posts/%v", domain, id)
	req, err := s.client.NewRequest("PATCH", u, post)
	if err != nil {
		return nil, nil, err
	}

	g := new(Post)
	resp, err := s.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Delete a post.
//
// Docbase API docs: https://help.docbase.io/posts/92982
func (s *PostService) Delete(ctx context.Context, domain string, id int64) (*Response, error) {
	u := fmt.Sprintf("teams/%v/posts/%v", domain, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
