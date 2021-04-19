package object

// Container is used by pod and deployment
type Container struct {
	Name            string         `json:"Name"`
	Image           string         `json:"Image"`
	ImagePullPolicy string         `json:"Image Pull Policy"`
	ContainerPorts  map[int]string `json:"Container Ports"`
}

// ObjectMeta is used by deployment, node, pod, service
type ObjectMeta struct {
	Name         string `json:"Name"`
	Namespace    string `json:"Namespace"`
	Uid          string `json:"UID"`
	CreationTime string `json:"Creation Time"`
}

type LabelSelector struct {
	MatchLabels      map[string]string          `json:"Match Labels"`
	MatchExpressions []LabelSelectorRequirement `json:"Match Expressions"`
}

type LabelSelectorRequirement struct {
	Key      string                `json:"Key"`
	Operator LabelSelectorOperator `json:"Operator"`
	Values   []string              `json:"Values"`
}

type LabelSelectorOperator string

const (
	LabelSelectorOpIn           LabelSelectorOperator = "In"
	LabelSelectorOpNotIn        LabelSelectorOperator = "NotIn"
	LabelSelectorOpExists       LabelSelectorOperator = "Exists"
	LabelSelectorOpDoesNotExist LabelSelectorOperator = "DoesNotExist"
)

// Role / ClusterRole / RoleBinding / ClusterRoleBinding
type PolicyRule struct {
	APIGroups       []string `json:"API Groups"`
	NonResourceURLs []string `json:"NonResource URLs"`
	ResourceNames   []string `json:"Resource Names"`
	Resources       []string `json:"Resources"`
	Verbs           []string `json:"Verbs"`
}

// Subject contains a reference to the object/user identities a Cluster/RoleBinding applies to
type Subject struct {
	Kind      string `json:"Kind"`
	APIGroup  string `json:"API Group"`
	Name      string `json:"Name"`
	Namespace string `json:"Namespace"`
}

// Subject contains a reference to the object/user identities a Cluster/RoleBinding applies to
type RoleRef struct {
	APIGroup string `json:"API Group"`
	Kind     string `json:"Kind"`
	Name     string `json:"Name"`
}
