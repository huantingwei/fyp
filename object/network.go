package object

type LoadBalancerIngress struct {
	HostName *string
	IP       *string
}

type LoadBalancerStatus struct {
	Ingress *LoadBalancerIngress
}

type Service struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Spec        *ServiceSpec
	Status      *ServiceStatus
}

type ServiceSpec struct {
	Type           *string
	Selector       map[string]string
	ClusterIP      *string
	ExternalIPs    []*string
	LoadBalancerIP *string
	Ports          []*ServicePort
}

type ServiceStatus struct {
	LoadBalancer *LoadBalancerStatus
}

type ServicePort struct {
	Name       *string
	Port       *int
	Protocol   *string
	NodePort   *int
	TargetPort interface{}
}

type Ingress struct {
	VersionKind *VersionKind
	Metadata    *ObjectMeta
	Spec        *IngressSpec
	Status      *IngressStatus
}

type IngressSpec struct {
	// DefaultBackend *IngressBackend
	IngressClassName *string
	Rules            []*IngressRule
	TLS              []*IngressTLS
}

type IngressStatus struct {
	LoadBalancer *LoadBalancerStatus
}

type IngressRule struct {
	Host *string
	Http interface{} // HTTPIngressRuleValue -> []HTTPIngressPath
}

type IngressTLS struct {
	Hosts      []*string
	SecretName *string
}
