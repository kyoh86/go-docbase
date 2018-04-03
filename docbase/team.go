package docbase

import (
	"context"
)

// TeamService provides access to the installation related functions
// in the Docbase API.
type TeamService service

// TeamListOptions specifies the optional parameters to the
// TeamService.List methods.
type TeamListOptions struct {
	ListOptions
}

// List teams.
//
// Docbase API docs: https://help.docbase.io/posts/92977
func (s *TeamService) List(ctx context.Context, opt *TeamListOptions) ([]*Team, *Response, error) {
	u := "teams"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var teams []*Team
	resp, err := s.client.Do(ctx, req, &teams)
	if err != nil {
		return nil, resp, err
	}

	return teams, resp, nil
}
