package object

// Service is a Pod with networking functionality
type Service struct {
	//Object meta
	ObjectMeta ObjectMeta `json:"objectMeta"`

	//Service spec
	ClusterIP      string            `json:"cluterIP"`
	ServiceType    string            `json:"serviceType"`
	LabelSelectors map[string]string `json:"labelSelectors"`
	ServicePorts   []ServicePort     `json:"servicePorts"`

	//load balancer status
	IngressIP []string `json:"ingressIP"`
}

// ServicePort is used by Service
type ServicePort struct {
	Port       int    `json:"port"`
	NodePort   int    `json:"nodePort"`
	TargetPort int    `json:"targetPort"`
	Protocol   string `json:"protocol"`
}
