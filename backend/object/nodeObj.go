package object

// Node is the vm instance
type Node struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"Object Meta"`

	//Node spec
	PodCIDR string `json:"Pod CIDR"`
	NodeID  string `json:"Node ID"`

	//Node system info
	MachineID        string `json:"Machine ID"`
	KernelVersion    string `json:"Kernel Version"`
	OsImage          string `json:"OS Image"`
	Os               string `json:"OS"`
	ContainerRuntime string `json:"Container Runtime"`
	KubeletVersion   string `json:"Kubelet Version"`
	KubeProxyVersion string `json:"KubeProxy Version"`

	//resources capacity
	CPUCap              int     `json:"CPU Cap"`
	MemoryCap           float64 `json:"Memory Cap"`
	PodsCap             int     `json:"Pods Cap"`
	EphemeralStorageCap float64 `json:"Ephemeral Storage Cap"`
	StorageCap          int     `json:"StorageCap"`

	//resources allocatable
	CPUAllocatable              int     `json:"CPU Allocatable"`
	MemoryAllocatable           float64 `json:"Memory Allocatable"`
	PodsAllocatable             int     `json:"Pods Allocatable"`
	EphemeralStorageAllocatable float64 `json:"Ephemeral Storage Allocatable"`
	StorageAllocatable          int     `json:"Storage Allocatable"`

	//Node conditions
	Conditions []Condition `json:"Conditions"`
}

// Condition of the node
type Condition struct {
	ConditionName      string `json:"Condition Name"`
	Status             string `json:"Status"` //true, false, or unknown
	LastHeartbeatTime  string `json:"Last Heartbeat Time"`
	LastTransitionTime string `json:"Last Transition Time"`
	Message            string `json:"Message"`
}
