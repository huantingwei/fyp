package object

// NetworkPolicy
type NetworkPolicy struct {
	ObjectMeta ObjectMeta 		`json:"objectMeta"`

	// NetworkPolicy spec
	Spec	NetworkPolicySpec	`json:"networkPolicySpec"`
}

// NetworkPolicySpec
type NetworkPolicySpec struct {
	NetworkPolicyEgressRule			[]NetworkPolicyEgressRule `json:"networkPolicyEgressRule"`
	NetworkPolicyIngressRule 		[]NetworkPolicyIngressRule `json:"networkPolicyIngressRule"`
	PodSelector						map[string]string `json:"podSelector"`
	PolicyTypes	 					[]PolicyType `json:"policyTypes"`
}

type NetworkPolicyEgressRule struct {
	Ports	[]NetworkPolicyPort `json:"ports"`
	To		[]NetworkPolicyPeer	`json:"to"`
}

type NetworkPolicyIngressRule struct {
	Ports	[]NetworkPolicyPort `json:"ports"`
	From	[]NetworkPolicyPeer	`json:"from"`
}

type NetworkPolicyPort struct {
	Port		interface{} `json:"port"`
	Protocol	Protocol		`json:"protocol"`
}

type NetworkPolicyPeer struct {
	// CIDR and Except belongs to NetworkPolicyPeer.IPBlock
	CIDR				string				`json:"cidr"`
	Except				string				`json:"except"`
	NamespaceSelector	map[string]string	`json:"namespaceSelector"`
	PodSelector			map[string]string	`json:"podspaceSelector"`
}

type PolicyType string
type Protocol string

