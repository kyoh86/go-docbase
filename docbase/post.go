package docbase

import (
	"context"
	"fmt"
	"time"
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
func (s *PostService) List(ctx context.Context, domain Domain, opt *PostListOptions) ([]Post, *Response, error) {
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
func (s *PostService) Get(ctx context.Context, domain Domain, id PostID) (*Post, *Response, error) {
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
	// Required.
	Title string `json:"title"`
	Body  string `json:"body"`

	// Option.
	Draft  bool      `json:"draft,omitempty"`
	Groups []GroupID `json:"groups,omitempty"`
	Scope  Scope     `json:"scope,omitempty"`
	Tags   []string  `json:"tags,omitempty"`
	Notice bool      `json:"notice,omitempty"`

	// Parameters that only owners can use.
	AuthorID    *UserID    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

// Create a post for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92980
func (s *PostService) Create(ctx context.Context, domain Domain, post *CreatePostOptions) (*Post, *Response, error) {
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
func (s *PostService) Edit(ctx context.Context, domain Domain, id PostID, post *PostEditRequest) (*Post, *Response, error) {
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

// PostEditRequest represents a request body for `Edit` function.
type PostEditRequest struct {
	Body   *string   `json:"body,omitempty"`
	Draft  *bool     `json:"draft,omitempty"`
	Groups []GroupID `json:"groups,omitempty"`
	Scope  *Scope    `json:"scope,omitempty"`
	Tags   []string  `json:"tags,omitempty"`
	Title  *string   `json:"title,omitempty"`
}

// Delete a post.
//
// Docbase API docs: https://help.docbase.io/posts/92982
func (s *PostService) Delete(ctx context.Context, domain Domain, id PostID) (*Response, error) {
	u := fmt.Sprintf("teams/%v/posts/%v", domain, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
