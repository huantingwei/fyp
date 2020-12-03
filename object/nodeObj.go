package object

type Node struct{
	//Object metadata
	ObjectMeta ObjectMeta

	//Node spec
	PodCIDR string
	NodeID string

	//Node system info
	MachineID string
	KernelVersion string
	OsImage string
	Os string
	ContainerRuntime string
	KubeletVersion string
	KubeProxyVersion string

	//resources capacity
	CpuCap int
	MemoryCap float64
	PodsCap int
	EphemeralStorageCap float64
	StorageCap int

	//resources allocatable
	CpuAllocatable int
	MemoryAllocatable float64
	PodsAllocatable int
	EphemeralStorageAllocatable float64
	StorageAllocatable int

	//Node conditions
	Conditions []Condition
}

type Condition struct{
	ConditionName string
	Status string //true, false, or unknown
	LastHeartbeatTime string
	LastTransitionTime string
	Message string
}