package object

// Service is a Pod with networking functionality
type Service struct {
	//Object meta
	ObjectMeta ObjectMeta `json:"Object Meta"`

	//Service spec
	ClusterIP      string            `json:"Cluter IP"`
	ServiceType    string            `json:"Service Type"`
	LabelSelectors map[string]string `json:"Label Selectors"`
	ServicePorts   []ServicePort     `json:"Service Ports"`

	//load balancer status
	IngressIP []string `json:"IngressIP"`
}

// ServicePort is used by Service
type ServicePort struct {
	Port       int    `json:"Port"`
	NodePort   int    `json:"Node Port"`
	TargetPort int    `json:"Target Port"`
	Protocol   string `json:"Protocol"`
}
