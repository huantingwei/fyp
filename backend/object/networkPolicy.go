package object

// NetworkPolicy
type NetworkPolicy struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// NetworkPolicy spec
	NetworkPolicyEgressRule  []NetworkPolicyEgressRule  `json:"NetworkPolicy Egress Rule"`
	NetworkPolicyIngressRule []NetworkPolicyIngressRule `json:"NetworkPolicy Ingress Rule"`
	PolicyTypes              []string                   `json:"Policy Types"`
}

type NetworkPolicyEgressRule struct {
	Ports []NetworkPolicyPort `json:"Ports"`
	To    []NetworkPolicyPeer `json:"To"`
}

type NetworkPolicyIngressRule struct {
	Ports []NetworkPolicyPort `json:"Ports"`
	From  []NetworkPolicyPeer `json:"From"`
}

type NetworkPolicyPort struct {
	Port     int         `json:"Port"`
	Protocol interface{} `json:"Protocol"`
}

type NetworkPolicyPeer struct {
	// CIDR and Except belongs to NetworkPolicyPeer.IPBlock
	CIDR              string            `json:"CIDR"`
	Except            []string          `json:"Except"`
	NamespaceSelector map[string]string `json:"Namespace Selector"`
	PodSelector       map[string]string `json:"Pod Selector"`
}
