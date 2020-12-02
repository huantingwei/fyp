type pod struct{
	//Object metadata
	podname string
	namespace string
	uid string
	creationTime string

	//object metadata - extra
	labels map[string]string
	//details of the pod owner, e.g. deployment
	ownerReferences []ownerReference 

	//Pod spec
	containers []container
	dnsPolicy string
	restartPolicy string
	nodeName string

	//Pod status
	hostIP string
	podIP string
	phase string //pending, running, failed, etc
	containerStatuses []containerStatus
}

type container struct{
	name string
	image string
	pullPolicy string
	containerPorts map[int]string
}

type ownerReference struct{
	name string
	uid string
	kind string
}

type containerStatus struct{
	name string
	image string
	ready bool
	restartCount int
}