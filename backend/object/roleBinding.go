package object

// RoleBinding
type RoleBinding struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`

	// RoleBinding spec
	Subjects	[]Subject	`json:"subjects"`
	RoleRef		RoleRef		`json:"roleRef"`
}
