package docbase

import (
	"context"
	"fmt"
)

// GroupService provides access to the installation related functions
// in the Docbase API.
type GroupService service

// GroupListOptions specifies the optional parameters to the
// GroupService.List methods.
type GroupListOptions struct {
	ListOptions
}

// List groups for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92978
func (s *GroupService) List(ctx context.Context, domain Domain, opt *GroupListOptions) ([]*Group, *Response, error) {
	u := fmt.Sprintf("teams/%v/groups", domain)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var groups []*Group
	resp, err := s.client.Do(ctx, req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}
