package docbase

// Post represents a Docbase Post.
type Post struct {
	ID            PostID    `json:"id"`
	Title         string    `json:"title"`
	Body          string    `json:"body"`
	Draft         bool      `json:"draft"`
	Archived      bool      `json:"archived"`
	URL           string    `json:"url"`
	CreatedAt     string    `json:"created_at"`
	Scope         Scope     `json:"scope"`
	SharingURL    string    `json:"sharing_url"`
	Tags          []Tag     `json:"tags"`
	User          User      `json:"user"`
	StarsCount    int64       `json:"stars_count"`
	GoodJobsCount int64       `json:"good_jobs_count"`
	Comments      []Comment `json:"comments"`
	Groups        []Group   `json:"groups"`
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
