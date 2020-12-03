package object

type Pod struct{
	//Object metadata
	ObjectMeta ObjectMeta

	//object metadata - extra
	Labels map[string]string
	//details of the pod owner, e.g. deployment
	OwnerReferences []OwnerReference 

	//Pod spec
	Containers []Container
	DnsPolicy string
	RestartPolicy string
	NodeName string

	//Pod status
	HostIP string
	PodIP string
	Phase string //pending, running, failed, etc
	ContainerStatuses []ContainerStatus
}

type OwnerReference struct{
	Name string
	Uid string
	Kind string
}

type ContainerStatus struct{
	Name string
	Image string
	Ready bool
	RestartCount int
}