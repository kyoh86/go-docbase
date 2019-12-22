package docbase

import (
	"context"
	"fmt"
)

// groupService provides access to the installation related functions
// in the Docbase API.
type groupService service

// Create a group for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/652985
func (s *groupService) Create(name string) *groupCreateDoer {
	return &groupCreateDoer{client: s.client, opts: groupCreateOptions{Name: name}}
}

// groupCreateOptions represents a Docbase Group.
type groupCreateOptions struct {
	// Required.
	Name string `json:"name"`

	// Option.
	Description *string `json:"description,omitempty"`
}

type groupCreateDoer struct {
	client *Client
	opts   groupCreateOptions
}

func (d *groupCreateDoer) Description(description string) *groupCreateDoer {
	d.opts.Description = &description
	return d
}

func (d *groupCreateDoer) Do(ctx context.Context) (*Group, *Response, error) {
	req, err := d.client.NewRequest("POST", "groups", d.opts)
	if err != nil {
		return nil, nil, err
	}

	g := new(Group)
	resp, err := d.client.Do(ctx, req, g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}

// List groups for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/92978
func (s *groupService) List() *groupListDoer {
	return &groupListDoer{client: s.client}
}

// groupListOptions specifies the parameters to the
// groupService.List methods.
type groupListOptions struct {
	// Option.
	Name *string `json:"name,omitempty"`

	ListOptions
}

type groupListDoer struct {
	client *Client
	opts   groupListOptions
}

func (d *groupListDoer) Page(page int64) *groupListDoer {
	d.opts.Page = &page
	return d
}

func (d *groupListDoer) PerPage(perPage int64) *groupListDoer {
	d.opts.PerPage = &perPage
	return d
}

func (d *groupListDoer) Do(ctx context.Context) ([]Group, *Response, error) {
	u, err := addOptions("groups", d.opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var groups []Group
	resp, err := d.client.Do(ctx, req, &groups)
	if err != nil {
		return nil, resp, err
	}

	return groups, resp, nil
}

// Get a single group.
//
// Docbase API docs: https://help.docbase.io/posts/652983
func (s *groupService) Get(id GroupID) *groupGetDoer {
	return &groupGetDoer{client: s.client, id: id}
}

type groupGetDoer struct {
	client *Client
	id     GroupID
}

func (d *groupGetDoer) Do(ctx context.Context) (*Group, *Response, error) {
	u := fmt.Sprintf("groups/%v", d.id)
	req, err := d.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	group := new(Group)
	resp, err := d.client.Do(ctx, req, group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// AddUsers a group for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/652985
func (s *groupService) AddUsers(groupID GroupID, userIDs []UserID) *groupAddUsersDoer {
	return &groupAddUsersDoer{client: s.client, groupID: groupID, opts: groupAddUsersOptions{UserIDs: userIDs}}
}

// groupAddUsersOptions represents a Docbase Group.
type groupAddUsersOptions struct {
	// Required.
	UserIDs []UserID `json:"user_ids"`
}

type groupAddUsersDoer struct {
	client  *Client
	groupID GroupID
	opts    groupAddUsersOptions
}

func (d *groupAddUsersDoer) Do(ctx context.Context) (*Response, error) {
	u := fmt.Sprintf("groups/%v/users", d.groupID)
	req, err := d.client.NewRequest("POST", u, d.opts)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// RemoveUsers a group for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/652985
func (s *groupService) RemoveUsers(groupID GroupID, userIDs []UserID) *groupRemoveUsersDoer {
	return &groupRemoveUsersDoer{client: s.client, groupID: groupID, opts: groupRemoveUsersOptions{UserIDs: userIDs}}
}

// groupRemoveUsersOptions represents a Docbase Group.
type groupRemoveUsersOptions struct {
	// Required.
	UserIDs []UserID `json:"user_ids"`
}

type groupRemoveUsersDoer struct {
	client  *Client
	groupID GroupID
	opts    groupRemoveUsersOptions
}

func (d *groupRemoveUsersDoer) Do(ctx context.Context) (*Response, error) {
	u := fmt.Sprintf("groups/%v/users", d.groupID)
	req, err := d.client.NewRequest("DELETE", u, d.opts)
	if err != nil {
		return nil, err
	}

	resp, err := d.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
