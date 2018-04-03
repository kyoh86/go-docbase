package docbase

// User represents a Docbase User.
type User struct {
	ID              UserID `json:"id"`
	Name            string `json:"name"`
	ProfileImageURL string `json:"profile_image_url"`
}

// UserID identifies a User.
type UserID int64
