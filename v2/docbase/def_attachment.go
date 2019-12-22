package docbase

import "time"

type Attachment struct {
	ID        AttachmentID `json:"id"`
	Name      string       `json:"name"`
	Size      int64        `json:"size"`
	URL       string       `json:"url"`
	Markdown  string       `json:"markdown"`
	CreatedAt time.Time    `json:"created_at"`
}

type AttachmentID int64
