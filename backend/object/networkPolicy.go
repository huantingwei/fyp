package object

// NetworkPolicy
type NetworkPolicy struct {
	ObjectMeta ObjectMeta 		`json:"objectMeta"`

	// NetworkPolicy spec
	NetworkPolicyEgressRule			[]NetworkPolicyEgressRule 	`json:"networkPolicyEgressRule"`
	NetworkPolicyIngressRule 		[]NetworkPolicyIngressRule 	`json:"networkPolicyIngressRule"`
	PolicyTypes	 					[]string					`json:"policyTypes"`
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
	Port		int 		`json:"port"`
	Protocol	interface{}	`json:"protocol"`
}

type NetworkPolicyPeer struct {
	// CIDR and Except belongs to NetworkPolicyPeer.IPBlock
	CIDR				string				`json:"cidr"`
	Except				[]string			`json:"except"`
	NamespaceSelector	map[string]string	`json:"namespaceSelector"`
	PodSelector			map[string]string	`json:"podspaceSelector"`
}


