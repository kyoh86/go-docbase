package docbase

// Post represents a Docbase Post.
type Post struct {
	Body      string    `json:"body"`
	Comments  []Comment `json:"comments"`
	CreatedAt string    `json:"created_at"`
	Draft     bool      `json:"draft"`
	Groups    []Group   `json:"groups"`
	ID        PostID    `json:"id"`
	Scope     Scope     `json:"scope"`
	Tags      []Tag     `json:"tags"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	User      User      `json:"user"`
}

// PostID specifies a post id for some API parameters.
type PostID int64

// Scope specifies a scope of the post.
type Scope string

const (
	// ScopeEveryone specifies that a post is published for everybody in the team.
	ScopeEveryone = Scope("everyone")
	// ScopeGroup specifies that a post is published for members in specified groups.
	ScopeGroup = Scope("group")
	// ScopePrivate specifies that a post is just for me.
	ScopePrivate = Scope("private")
)

func (s Scope) String() string { return string(s) }
