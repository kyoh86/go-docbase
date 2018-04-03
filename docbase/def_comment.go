package docbase

// Comment represents a Docbase Comment.
type Comment struct {
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
	ID        int64  `json:"id"`
	User      User   `json:"user"`
}
