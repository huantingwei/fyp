package object

// Role
type Role struct {
	ObjectMeta ObjectMeta `json:"objectMeta"`

	// Role spec
	Rules      []PolicyRule	`json:"rules"`
}

type PolicyRule struct {
	APIGroups 		[]string `json:"apiGroups"`
	NonResourceURLs []string `json:"nonResourceUrls"`
	ResourceNames 	[]string `json:"resourceNames"`
	Resources 		[]string `json:"resources"`
	Verbs 			[]string `json:"verbs"`
}