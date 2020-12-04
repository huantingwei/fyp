package object

type Deployment struct{
	//Object metadata
	ObjectMeta ObjectMeta

	//deployment spec
	DesiredPods int
	MatchLabels map[string]string

	//Pod spec for this deployment
	Containers []Container
	DnsPolicy string
	RestartPolicy string
	NodeName string

	//deployment status
	UpdatedReplicas int
	ReadyReplicas int
	AvailableReplicas int
	UnavailableReplicas int
	
}

