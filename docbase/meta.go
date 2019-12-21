package docbase

import (
	"net/url"
	"strconv"
)

// Meta provides the page values for paginating through a set of
// results. Any or all of these may be set to the zero value for
// responses that are not part of a paginated set, or for which there
// are no additional pages.
type Meta struct {
	NextPage     string `json:"next_page"`
	PreviousPage string `json:"previous_page"`
	Total        int64  `json:"total"`
}

// Previous will get a ListOptions for the previous page.
// If there's no previous page, it will be nil.
func (m Meta) Previous() *ListOptions {
	if m.PreviousPage == "" {
		return nil
	}
	u, err := url.Parse(m.PreviousPage)
	if err != nil {
		return nil
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}
	n, err := strconv.ParseInt(q.Get("page"), 10, 64)
	if err != nil {
		return nil
	}
	p, err := strconv.ParseInt(q.Get("per_page"), 10, 64)
	if err != nil {
		return nil
	}
	return &ListOptions{
		Page:    &n,
		PerPage: &p,
	}
}

// Next will get a ListOptions for the next page.
// If there's no next page, it will be nil.
func (m Meta) Next() *ListOptions {
	if m.NextPage == "" {
		return nil
	}
	u, err := url.Parse(m.NextPage)
	if err != nil {
		return nil
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil
	}
	n, err := strconv.ParseInt(q.Get("page"), 10, 64)
	if err != nil {
		return nil
	}
	p, err := strconv.ParseInt(q.Get("per_page"), 10, 64)
	if err != nil {
		return nil
	}
	return &ListOptions{
		Page:    &n,
		PerPage: &p,
	}
}
