type service struct{
	//Object meta
	nodename string
	namespace string
	uid string
	creationTime string

	//Service spec
	clusterIP string
	serviceType string
	labelSelectors map[string]string
	servicePorts []servicePort

	//load balancer status
	ingressIP string

}

type servicePort struct{
	port int
	nodePort int
	targetPort int
	protocol string
}