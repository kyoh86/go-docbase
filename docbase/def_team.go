package docbase

// Team represents a Docbase Team.
type Team struct {
	Domain Domain `json:"domain"`
	Name   string `json:"name"`
}
