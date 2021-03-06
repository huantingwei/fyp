package object

// Node is the vm instance
type Node struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"objectMeta"`

	//Node spec
	PodCIDR string `json:"podCIDR"`
	NodeID  string `json:"nodeID"`

	//Node system info
	MachineID        string `json:"machineID"`
	KernelVersion    string `json:"kernelVersion"`
	OsImage          string `json:"osImage"`
	Os               string `json:"os"`
	ContainerRuntime string `json:"containerRuntime"`
	KubeletVersion   string `json:"kubeletVersion"`
	KubeProxyVersion string `json:"kubeProxyVersion"`

	//resources capacity
	CPUCap              int     `json:"cpuCap"`
	MemoryCap           float64 `json:"memoryCap"`
	PodsCap             int     `json:"podsCap"`
	EphemeralStorageCap float64 `json:"ephemeralStorageCap"`
	StorageCap          int     `json:"storageCap"`

	//resources allocatable
	CPUAllocatable              int     `json:"cpuAllocatable"`
	MemoryAllocatable           float64 `json:"memoryAllocatable"`
	PodsAllocatable             int     `json:"podsAllocatable"`
	EphemeralStorageAllocatable float64 `json:"ephemeralStorageAllocatable"`
	StorageAllocatable          int     `json:"storageAllocatable"`

	//Node conditions
	Conditions []Condition `json:"conditions"`
}

// Condition of the node
type Condition struct {
	ConditionName      string `json:"conditionName"`
	Status             string `json:"status"` //true, false, or unknown
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
}
