package docbase

import (
	"context"
	"fmt"
	"time"
)

// postService handles communication with the Post related
// methods of the Docbase API.
type postService service

// List posts in a domain.
// To build query, use docbase/postquery package.
//
// Docbase API docs: https://help.docbase.io/posts/92984
// Docbase query docs: https://help.docbase.io/posts/59432
func (s *postService) List() *postListDoer {
	return &postListDoer{client: s.client}
}

// postListOptions specifies the optional parameters to the
// postService.List method.
type postListOptions struct {
	// Query filters Post.
	//
	// docs: https://help.docbase.io/posts/59432?list=%2Fsearch&q=%E6%A4%9C%E7%B4%A2#%E6%A4%9C%E7%B4%A2%E3%82%AA%E3%83%97%E3%82%B7%E3%83%A7%E3%83%B3
	Query *string `url:"q,omitempty"`

	ListOptions
}

type postListDoer struct {
	client *Client
	opts   postListOptions
}

func (d *postListDoer) Query(query string) *postListDoer {
	d.opts.Query = &query
	return d
}

func (d *postListDoer) Page(page int64) *postListDoer {
	d.opts.Page = &page
	return d
}

func (d *postListDoer) PerPage(perPage int64) *postListDoer {
	d.opts.PerPage = &perPage
	return d
}

func (d *postListDoer) Do(ctx context.Context) ([]Post, *Response, error) {
	u, err := addOptions("posts", d.opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var posts struct {
		Meta  Meta   `json:"meta"`
		Posts []Post `json:"posts"`
	}
	resp, err := d.client.Do(ctx, req, &posts)
	if err != nil {
		return nil, resp, err
	}
	resp.Meta = posts.Meta

	return posts.Posts, resp, nil
}

// Get a single post.
//
// Docbase API docs: https://help.docbase.io/posts/97204
func (s *postService) Get(id PostID) *postGetDoer {
	return &postGetDoer{client: s.client, id: id}
}

type postGetDoer struct {
	client *Client
	id     PostID
}

func (d *postGetDoer) Do(ctx context.Context) (*Post, *Response, error) {
	u := fmt.Sprintf("posts/%v", d.id)
	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	post := new(Post)
	resp, err := d.client.Do(ctx, req, post)
	if err != nil {
		return nil, resp, err
	}

	return post, resp, nil
}

// Create a post for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92980
func (s *postService) Create(title, body string) *postCreateDoer {
	return &postCreateDoer{client: s.client, opts: postCreateOptions{Title: title, Body: body}}
}

// postCreateOptions represents a Docbase Post.
type postCreateOptions struct {
	// Required.
	Title string `json:"title"`
	Body  string `json:"body"`

	// Option.
	Draft  *bool      `json:"draft,omitempty"`
	Groups *[]GroupID `json:"groups,omitempty"`
	Scope  *Scope     `json:"scope,omitempty"`
	Tags   *[]string  `json:"tags,omitempty"`
	Notice *bool      `json:"notice,omitempty"`

	// Parameters that only owners can use.
	AuthorID    *UserID    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type postCreateDoer struct {
	client *Client
	opts   postCreateOptions
}

func (d *postCreateDoer) Draft(draft bool) *postCreateDoer {
	d.opts.Draft = &draft
	return d
}

func (d *postCreateDoer) Groups(groups []GroupID) *postCreateDoer {
	d.opts.Groups = &groups
	return d
}

func (d *postCreateDoer) Scope(scope Scope) *postCreateDoer {
	d.opts.Scope = &scope
	return d
}

func (d *postCreateDoer) Tags(tags []string) *postCreateDoer {
	d.opts.Tags = &tags
	return d
}

func (d *postCreateDoer) Notice(notice bool) *postCreateDoer {
	d.opts.Notice = &notice
	return d
}

func (d *postCreateDoer) AuthorID(authorID UserID) *postCreateDoer {
	d.opts.AuthorID = &authorID
	return d
}

func (d *postCreateDoer) PublishedAt(publishedAt time.Time) *postCreateDoer {
	d.opts.PublishedAt = &publishedAt
	return d
}

func (d *postCreateDoer) Do(ctx context.Context) (*Post, *Response, error) {
	req, err := d.client.NewRequest("POST", "posts", d.opts)
	if err != nil {
		return nil, nil, err
	}

	g := new(Post)
	resp, err := d.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Edit a post.
//
// Docbase API docs: https://help.docbase.io/posts/92981
func (s *postService) Edit(id PostID) *postEditDoer {
	return &postEditDoer{client: s.client, id: id}
}

// postEditOptions represents a request body for `Edit` function.
type postEditOptions struct {
	Body   *string    `json:"body,omitempty"`
	Notice *bool      `json:"notice,omitempty"`
	Draft  *bool      `json:"draft,omitempty"`
	Groups *[]GroupID `json:"groups,omitempty"`
	Scope  *Scope     `json:"scope,omitempty"`
	Tags   *[]string  `json:"tags,omitempty"`
	Title  *string    `json:"title,omitempty"`
}

type postEditDoer struct {
	client *Client
	id     PostID
	opts   postEditOptions
}

func (d *postEditDoer) Body(body string) *postEditDoer {
	d.opts.Body = &body
	return d
}

func (d *postEditDoer) Notice(notice bool) *postEditDoer {
	d.opts.Notice = &notice
	return d
}

func (d *postEditDoer) Draft(draft bool) *postEditDoer {
	d.opts.Draft = &draft
	return d
}

func (d *postEditDoer) Groups(groups []GroupID) *postEditDoer {
	d.opts.Groups = &groups
	return d
}

func (d *postEditDoer) Scope(scope Scope) *postEditDoer {
	d.opts.Scope = &scope
	return d
}

func (d *postEditDoer) Tags(tags []string) *postEditDoer {
	d.opts.Tags = &tags
	return d
}

func (d *postEditDoer) Title(title string) *postEditDoer {
	d.opts.Title = &title
	return d
}

func (d *postEditDoer) Do(ctx context.Context) (*Post, *Response, error) {
	u := fmt.Sprintf("posts/%v", d.id)
	req, err := d.client.NewRequest("PATCH", u, d.opts)
	if err != nil {
		return nil, nil, err
	}

	g := new(Post)
	resp, err := d.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Archive a single post.
//
// Docbase API docs: https://help.docbase.io/posts/97204
func (s *postService) Archive(id PostID) *postArchiveDoer {
	return &postArchiveDoer{client: s.client, id: id}
}

type postArchiveDoer struct {
	client *Client
	id     PostID
}

func (d *postArchiveDoer) Do(ctx context.Context) (*Response, error) {
	u := fmt.Sprintf("posts/%v/archive", d.id)
	req, err := d.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(ctx, req, nil)
}

// Unarchive a single post.
//
// Docbase API docs: https://help.docbase.io/posts/97204
func (s *postService) Unarchive(id PostID) *postUnarchiveDoer {
	return &postUnarchiveDoer{client: s.client, id: id}
}

type postUnarchiveDoer struct {
	client *Client
	id     PostID
}

func (d *postUnarchiveDoer) Do(ctx context.Context) (*Response, error) {
	u := fmt.Sprintf("posts/%v/unarchive", d.id)
	req, err := d.client.NewRequest("PUT", u, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(ctx, req, nil)
}
