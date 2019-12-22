package docbase

import (
	"context"
	"fmt"
	"time"
)

// commentService handles communication with the Comment related
// methods of the Docbase API.
type commentService service

// Create a comment for a post.
//
// Docbase API docs: https://help.docbase.io/posts/216289
func (s *commentService) Create(postID PostID, body string) *commentCreateDoer {
	return &commentCreateDoer{client: s.client, postID: postID, opts: commentCreateOptions{Body: body}}
}

// commentCreateOptions specifies the optional parameters to the
// commentService.Create methods.
type commentCreateOptions struct {
	Body   string `json:"body"`
	Notice *bool  `json:"notice,omitempty"`

	// Parameters that only owners can use.
	AuthorID    *UserID    `json:"author_id,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

type commentCreateDoer struct {
	client *Client
	postID PostID
	opts   commentCreateOptions
}

func (d *commentCreateDoer) Notice(notice bool) *commentCreateDoer {
	d.opts.Notice = &notice
	return d
}

func (d *commentCreateDoer) AuthorID(authorID UserID) *commentCreateDoer {
	d.opts.AuthorID = &authorID
	return d
}

func (d *commentCreateDoer) PublishedAt(publishedAt time.Time) *commentCreateDoer {
	d.opts.PublishedAt = &publishedAt
	return d
}

func (d *commentCreateDoer) Do(ctx context.Context) (*Comment, *Response, error) {
	u := fmt.Sprintf("posts/%v/comments", d.postID)
	req, err := d.client.NewRequest("POST", u, d.opts)
	if err != nil {
		return nil, nil, err
	}

	g := new(Comment)
	resp, err := d.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Delete a comment.
//
// Docbase API docs: https://help.docbase.io/posts/216290
func (s *commentService) Delete(id CommentID) *commentDeleteDoer {
	return &commentDeleteDoer{client: s.client, id: id}
}

type commentDeleteDoer struct {
	id     CommentID
	client *Client
}

func (d *commentDeleteDoer) Do(ctx context.Context) (*Response, error) {
	u := fmt.Sprintf("comments/%v", d.id)
	req, err := d.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return d.client.Do(ctx, req, nil)
}
