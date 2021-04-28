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

type NetworkPolicyV2 struct {
	ObjectMeta ObjectMeta `json:"Object Meta"`

	// NetworkPolicy spec
	NetworkPolicyEgressRule  []NetworkPolicyEgressRuleV2  `json:"NetworkPolicy Egress Rule"`
	NetworkPolicyIngressRule []NetworkPolicyIngressRuleV2 `json:"NetworkPolicy Ingress Rule"`
	PolicyTypes              string
}
type NetworkPolicyEgressRuleV2 struct {
	Ports map[string]interface{} `json:"Ports"` // protocol:port
	/*
		{
			"CIDR": "CIDR1, CIDR2, ...",
			"Except": "e1, e2, ...",
			"NamespaceSelector": "nl1:nv1, nl2:nv2, ...",
			"PodSelector": "pl1:pv1, pl2:pv2 ..."
		}
	*/
	To map[string]string `json:"To"`
}

type NetworkPolicyIngressRuleV2 struct {
	Ports map[string]interface{} `json:"Ports"`
	/*
		{
			"CIDR": "CIDR1, CIDR2, ...",
			"Except": "e1, e2, ...",
			"NamespaceSelector": "nl1:nv1, nl2:nv2, ...",
			"PodSelector": "pl1:pv1, pl2:pv2 ..."
		}
	*/
	From map[string]string `json:"From"`
}
