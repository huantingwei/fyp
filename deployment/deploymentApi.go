type deployment struct{
	//Object metadata
	deploymentname string
	namespace string
	uid string
	creationTime string

	//deployment spec
	desiredPods int
	matchLabels map[string]string

	//Pod spec for this deployment
	containers []container
	dnsPolicy string
	restartPolicy string

	//deployment status
	updatedReplicas int
	readyReplicas int
	availableReplicas int
	unavailableReplicas int
	
}

type container struct{
	name string
	image string
	pullPolicy string
	containerPorts map[int]string
}