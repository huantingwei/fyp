//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#container-v1-core
type container struct{
	image     string
	name      string
	resources resourceRequirements
	ports     containerPorts //which port to expose on the pod's IP address
}

//https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#resourcerequirements-v1-core
type resourceRequirements struct{
	limits   map[string]string //e.g. cpu, memory
	requests map[string]string //e.g. cpu, memory
}

//API doc: https://v1-16.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#containerport-v1-core
type containerPorts{
	containerPort int
	protocol      string //UDP, TCP, or SCTP
}