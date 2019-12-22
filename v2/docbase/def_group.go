package docbase

import "time"

// Group represents a Docbase Group.
type Group struct {
	ID             GroupID   `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	PostsCount     int64     `json:"posts_count"`
	LastActivityAt time.Time `json:"last_activity_at"`
	CreatedAt      time.Time `json:"created_at"`
	Users          []User
}

// GroupID identifies a Group.
type GroupID int64
