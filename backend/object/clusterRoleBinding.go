package object

// ClusterRoleBinding
type ClusterRoleBinding struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`

	// ClusterRoleBinding spec
	Subjects	[]Subject	`json:"subjects"`
	RoleRef		RoleRef		`json:"roleRef"`
}

