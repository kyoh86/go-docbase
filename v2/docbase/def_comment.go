package docbase

// Comment represents a Docbase Comment.
type Comment struct {
	Body      string    `json:"body"`
	CreatedAt string    `json:"created_at"`
	ID        CommentID `json:"id"`
	User      User      `json:"user"`
}

// CommentID identifies a Comment.
type CommentID int64
