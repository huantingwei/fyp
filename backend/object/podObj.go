package object

// Pod is the minimum execution unit
type Pod struct {
	//Object metadata
	ObjectMeta ObjectMeta `json:"objectMeta"`

	//object metadata - extra
	Labels map[string]string `json:"labels"`
	//details of the pod owner, e.g. deployment
	OwnerReferences []OwnerReference `json:"ownerReferences"`

	//Pod spec
	Containers    []Container `json:"containers"`
	DNSPolicy     string      `json:"DNSPolicy"`
	RestartPolicy string      `json:"restartPolicy"`
	NodeName      string      `json:"nodeName"`

	//Pod status
	HostIP            string            `json:"hostIP"`
	PodIP             string            `json:"podIP"`
	Phase             string            `json:"phase"` //pending, running, failed, etc
	ContainerStatuses []ContainerStatus `json:"containerStatuses"`
}

// OwnerReference specify what owns the Pod
type OwnerReference struct {
	Name string `json:"name"`
	Uid  string `json:"uid"`
	Kind string `json:"kind"`
}

// ContainerStatus describes the status of containers
type ContainerStatus struct {
	Name         string `json:"name"`
	Image        string `json:"image"`
	Ready        bool   `json:"ready"`
	RestartCount int    `json:"restartCount"`
}
