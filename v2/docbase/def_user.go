package docbase

import "time"

// User represents a Docbase User.
type User struct {
	ID                    UserID    `json:"id"`
	Name                  string    `json:"name"`
	Username              string    `json:"username"`
	ProfileImageURL       string    `json:"profile_image_url"`
	Role                  UserRole  `json:"role"`
	PostsCount            int64     `json:"posts_count"`
	LastAccessTime        time.Time `json:"last_access_time"`
	TwoStepAuthentication bool      `json:"two_step_authentication"`
	Groups                []Group   `json:"groups"`
}

// UserID identifies a User.
type UserID int64
type UserRole int64

const (
	UserRoleUser UserRole = iota
	UserRoleAdmin
	UserRoleOwner
)
