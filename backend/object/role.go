package object

// Role
type Role struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`

	// Role spec
	Rules      []PolicyRule	`json:"rules"`
}