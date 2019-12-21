package docbase

import (
	"context"
)

// tagService provides access to the installation related functions
// in the Docbase API.
type tagService service

// List tags for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92979
func (s *tagService) List() *tagListDoer {
	return &tagListDoer{client: s.client}
}

type tagListDoer struct {
	client *Client
}

func (d *tagListDoer) Do(ctx context.Context) ([]Tag, *Response, error) {
	req, err := d.client.NewRequest("GET", "tags", nil)
	if err != nil {
		return nil, nil, err
	}

	var tags []Tag
	resp, err := d.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, resp, err
	}

	return tags, resp, nil
}
