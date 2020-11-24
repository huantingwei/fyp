package object

type NetworkPolicy struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Spec        *NetworkPolicySpec
}

type NetworkPolicySpec struct {
	Egress      []*NetworkPolicyEgressRule
	Ingress     []*NetworkPolicyIngressRule
	PodSelector *LabelSelector
	PolicyTypes []*string
}

type NetworkPolicyEgressRule struct {
	Ports []*NetworkPolicyPort
	To    []*NetworkPolicyPeer
}

type NetworkPolicyIngressRule struct {
	Ports []*NetworkPolicyPort
	From  []*NetworkPolicyPeer
}

type NetworkPolicyPort struct {
	Port     interface{} // either *string or *int
	Protocol *string
}

type NetworkPolicyPeer struct {
	IpBlock           *IPBlock
	NameSpaceSelector *LabelSelector
	PodSelector       *LabelSelector
}

type IPBlock struct {
	Cidr   *string
	Except []*string
}
