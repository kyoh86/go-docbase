package docbase

// Post represents a Docbase Post.
type Post struct {
	Body      string    `json:"body"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"created_at"`
	Draft     bool      `json:"draft"`
	Groups    []Group   `json:"groups"`
	ID        PostID    `json:"id"`
	Scope     string    `json:"scope"`
	Tags      []Tag     `json:"tags"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	User      User      `json:"user"`
}

// PostID specifies a post id for some API parameters.
type PostID int64
