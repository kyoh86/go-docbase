package docbase

import (
	"context"
)

// attachmentService provides access to the installation related functions
// in the Docbase API.
type attachmentService service

// Upload attachments for a domain.
//
// Docbase API docs: https://help.docbase.io/posts/225804
func (s *attachmentService) Upload() *attachmentUploadDoer {
	return &attachmentUploadDoer{client: s.client}
}

type attachmentUploadPayload struct {
	Name    string `json:"name"`
	Content []byte `json:"content"`
}

type attachmentUploadDoer struct {
	client   *Client
	payloads []attachmentUploadPayload
}

func (d *attachmentUploadDoer) AddPayload(name string, content []byte) *attachmentUploadDoer {
	d.payloads = append(d.payloads, attachmentUploadPayload{Name: name, Content: content})
	return d
}

func (d *attachmentUploadDoer) Do(ctx context.Context) ([]Attachment, *Response, error) {
	req, err := d.client.NewRequest("POST", "attachments", d.payloads)
	if err != nil {
		return nil, nil, err
	}

	var g []Attachment
	resp, err := d.client.Do(ctx, req, &g)
	if err != nil {
		return nil, resp, err
	}

	return g, resp, nil
}
