type node struct{
	//Object meta
	nodename string
	namespace string
	uid string
	creationTime string

	//Node spec
	podCIDR string
	nodeID string

	//Node system info
	machineID string
	kernelVersion string
	osImage string
	os string
	containerRuntime string
	kubeletVersion string
	kubeProxyVersion string

	//resources capacity
	cpuCap int
	memoryCap float64
	podsCap int
	ephemeralStorageCap float64
	storageCap int

	//resources allocatable
	cpuAllocatable int
	memoryAllocatable float64
	podsAllocatable int
	ephemeralStorageAllocatable float64
	storageAllocatable int

	//Node conditions
	conditions []condition
}

type condition struct{
	conditionName string
	status string //true, false, or unknown
	lastHeartbeatTime string
	lastTransitionTime string
	message string
}