package object

// RoleBinding
type RoleBinding struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`

	// RoleBinding spec
	Subjects	[]Subject	`json:"subjects"`
	RoleRef		RoleRef		`json:"roleRef"`
}

// Subject contains a reference to the object/user identities a RoleBinding applies to
type Subject struct {
	Kind			string `json:"kind"`
	APIGroup 		string `json:"apiGroup"`
	Name			string `json:"name"`
	Namespace	 	string `json:"namespace"`
}

// Subject contains a reference to the object/user identities a RoleBinding applies to
type RoleRef struct {
	APIGroup 		string `json:"apiGroup"`
	Kind			string `json:"kind"`
	Name			string `json:"name"`
}