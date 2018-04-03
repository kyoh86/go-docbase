package docbase

import (
	"context"
	"fmt"
)

// CommentService handles communication with the Comment related
// methods of the Docbase API.
type CommentService service

// CommentListOptions specifies the optional parameters to the
// CommentService.List methods.
type CommentListOptions struct {
	ListOptions
}

// Create a comment for a post.
//
// Docbase API docs: https://help.docbase.io/posts/216289
func (s *CommentService) Create(ctx context.Context, domain Domain, postID PostID, comment *Comment) (*Comment, *Response, error) {
	u := fmt.Sprintf("teams/%v/posts/%v/comments", domain, postID)
	req, err := s.client.NewRequest("POST", u, comment)
	if err != nil {
		return nil, nil, err
	}

	g := new(Comment)
	resp, err := s.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// Delete a comment.
//
// Docbase API docs: https://help.docbase.io/posts/216290
func (s *CommentService) Delete(ctx context.Context, domain Domain, id PostID) (*Response, error) {
	u := fmt.Sprintf("teams/%v/comments/%v", domain, id)
	req, err := s.client.NewRequest("DELETE", u, nil)
	if err != nil {
		return nil, err
	}
	return s.client.Do(ctx, req, nil)
}
