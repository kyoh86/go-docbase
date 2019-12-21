package docbase

import (
	"context"
)

// userService handles communication with the User related
// methods of the Docbase API.
type userService service

// List users in a domain.
//
// Docbase API docs: https://help.docbase.io/posts/680809
// NOTE: In the document, "role" property is string but response has number.
//       Now this library implements based on reality. (22 Dec, 2019)
func (s *userService) List() *userListDoer {
	return &userListDoer{client: s.client}
}

// userListOptions specifies the optional parameters to the
// userService.List method.
type userListOptions struct {
	Query             *string `url:"q,omitempty"`
	IncludeUserGroups *bool   `url:"include_user_groups,omitempty"`

	ListOptions
}

type userListDoer struct {
	opts   userListOptions
	client *Client
}

func (d *userListDoer) Query(query string) *userListDoer {
	d.opts.Query = &query
	return d
}

func (d *userListDoer) IncludeUserGroups(includeUserGroups bool) *userListDoer {
	d.opts.IncludeUserGroups = &includeUserGroups
	return d
}

func (d *userListDoer) Page(page int64) *userListDoer {
	d.opts.Page = &page
	return d
}

func (d *userListDoer) PerPage(perPage int64) *userListDoer {
	d.opts.PerPage = &perPage
	return d
}

func (d *userListDoer) Do(ctx context.Context) ([]User, *Response, error) {
	u, err := addOptions("users", d.opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var users []User
	resp, err := d.client.Do(ctx, req, &users)
	if err != nil {
		return nil, resp, err
	}

	return users, resp, nil
}
