package object

// RoleBinding
type RoleBinding struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// RoleBinding spec
	Subjects []Subject `json:"Subjects"`
	RoleRef  RoleRef   `json:"Role Ref"`
}
