package object

// Container is used by pod and deployment
type Container struct {
	Name            string         `json:"name"`
	Image           string         `json:"image"`
	ImagePullPolicy string         `json:"imagePullPolicy"`
	ContainerPorts  map[int]string `json:"containerPorts"`
}

// ObjectMeta is used by deployment, node, pod, service
type ObjectMeta struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Uid          string `json:"uid"`
	CreationTime string `json:"creationTime"`
}

type LabelSelector struct {
	MatchLabels 		map[string]string 			`json:"matchLabels"`
	MatchExpressions 	[]LabelSelectorRequirement 	`json:"matchExpressions"`
}

type LabelSelectorRequirement struct {
	Key			string					`json:"key"`
	Operator	LabelSelectorOperator	`json:"operator"`
	Values		[]string				`json:"values"`
}

type LabelSelectorOperator string
const (
	LabelSelectorOpIn			LabelSelectorOperator  	= "In"
	LabelSelectorOpNotIn		LabelSelectorOperator	= "NotIn"
	LabelSelectorOpExists		LabelSelectorOperator	= "Exists"
	LabelSelectorOpDoesNotExist	LabelSelectorOperator	= "DoesNotExist"
)

// Role / ClusterRole / RoleBinding / ClusterRoleBinding
type PolicyRule struct {
	APIGroups 		[]string `json:"apiGroups"`
	NonResourceURLs []string `json:"nonResourceUrls"`
	ResourceNames 	[]string `json:"resourceNames"`
	Resources 		[]string `json:"resources"`
	Verbs 			[]string `json:"verbs"`
}

// Subject contains a reference to the object/user identities a Cluster/RoleBinding applies to
type Subject struct {
	Kind			string `json:"kind"`
	APIGroup 		string `json:"apiGroup"`
	Name			string `json:"name"`
	Namespace	 	string `json:"namespace"`
}

// Subject contains a reference to the object/user identities a Cluster/RoleBinding applies to
type RoleRef struct {
	APIGroup 		string `json:"apiGroup"`
	Kind			string `json:"kind"`
	Name			string `json:"name"`
}