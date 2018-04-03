package docbase

import (
	"context"
	"fmt"
)

// TagService provides access to the installation related functions
// in the Docbase API.
type TagService service

// TagListOptions specifies the optional parameters to the
// TagService.List methods.
type TagListOptions struct {
	ListOptions
}

// List tags for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92979
func (s *TagService) List(ctx context.Context, domain string, opt *TagListOptions) ([]*Tag, *Response, error) {
	u := fmt.Sprintf("teams/%v/tags", domain)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var tags []*Tag
	resp, err := s.client.Do(ctx, req, &tags)
	if err != nil {
		return nil, resp, err
	}

	return tags, resp, nil
}
