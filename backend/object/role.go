package object

// Role
type Role struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// Role spec
	Rules []PolicyRule `json:"Rules"`
}
