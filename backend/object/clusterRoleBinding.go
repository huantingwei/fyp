package object

// ClusterRoleBinding
type ClusterRoleBinding struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// ClusterRoleBinding spec
	Subjects []Subject `json:"Subjects"`
	RoleRef  RoleRef   `json:"Role Ref"`
}
