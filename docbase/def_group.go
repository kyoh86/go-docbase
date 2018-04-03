package docbase

// Group represents a Docbase Group.
type Group struct {
	ID   GroupID `json:"id"`
	Name string  `json:"name"`
}

// GroupID identifies a Group.
type GroupID int64
