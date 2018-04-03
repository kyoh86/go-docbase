package docbase

// User represents a Docbase User.
type User struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}
