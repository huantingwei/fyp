package object

type Service struct{
	//Object meta
	ObjectMeta ObjectMeta

	//Service spec
	ClusterIP string
	ServiceType string
	LabelSelectors map[string]string
	ServicePorts []ServicePort

	//load balancer status
	IngressIP string

}

type ServicePort struct{
	Port int
	NodePort int
	TargetPort int
	Protocol string
}